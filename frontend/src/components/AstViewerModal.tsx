import { useMemo } from 'react'
import { Modal } from './Modal'
import { useTranslation } from '../i18n/I18nContext'
import './AstViewerModal.css'

interface AstViewerModalProps {
  astText: string
  onClose: () => void
}

interface AstAttribute {
  label: string
  value: string
}

interface AstGraphNode {
  id: string
  type: string
  edgeLabel?: string
  attrs: AstAttribute[]
  children: AstGraphNode[]
}

type ParseNodeKind = 'nt' | 't' | 'eps'

interface ParseNode {
  id: string
  label: string
  kind: ParseNodeKind
  children: ParseNode[]
}

interface ParseLayoutNode extends ParseNode {
  children: ParseLayoutNode[]
  width: number
  height: number
  subtreeWidth: number
  x: number
  y: number
}

interface ParseEdge {
  id: string
  fromX: number
  fromY: number
  toX: number
  toY: number
}

interface ParseLayoutResult {
  width: number
  height: number
  nodes: ParseLayoutNode[]
  edges: ParseEdge[]
}

const CHAR_WIDTH = 8
const NODE_HEIGHT = 22
const LEVEL_GAP = 56
const SIBLING_GAP = 26
const DIAGRAM_PADDING_X = 36
const DIAGRAM_PADDING_Y = 26

function isNodeType(value: string): boolean {
  return /^[A-Za-z][A-Za-z0-9]*Node$/.test(value.trim())
}

function parseAst(astText: string): AstGraphNode | null {
  const lines = astText
    .split('\n')
    .map(line => line.replace(/\r/g, ''))
    .filter(line => line.trim() !== '')

  if (lines.length === 0) {
    return null
  }

  let nextId = 1
  const root: AstGraphNode = {
    id: `ast-${nextId++}`,
    type: lines[0].trim(),
    attrs: [],
    children: [],
  }
  const parentsByDepth: AstGraphNode[] = [root]

  for (const rawLine of lines.slice(1)) {
    const branchIndex = rawLine.indexOf('├── ') >= 0 ? rawLine.indexOf('├── ') : rawLine.indexOf('└── ')
    if (branchIndex < 0) {
      continue
    }

    const depth = Math.floor(branchIndex / 4)
    const content = rawLine.slice(branchIndex + 4).trim()
    const parent = parentsByDepth[depth]
    if (!parent) {
      continue
    }

    const pair = content.match(/^([^:]+):\s*(.+)$/)
    if (pair) {
      const label = pair[1].trim()
      const value = pair[2].trim()

      if (isNodeType(value)) {
        const child: AstGraphNode = {
          id: `ast-${nextId++}`,
          type: value,
          edgeLabel: label,
          attrs: [],
          children: [],
        }
        parent.children.push(child)
        parentsByDepth[depth + 1] = child
        parentsByDepth.length = depth + 2
        continue
      }

      parent.attrs.push({ label, value })
      parentsByDepth.length = depth + 1
      continue
    }

    if (!isNodeType(content)) {
      continue
    }

    const child: AstGraphNode = {
      id: `ast-${nextId++}`,
      type: content,
      attrs: [],
      children: [],
    }
    parent.children.push(child)
    parentsByDepth[depth + 1] = child
    parentsByDepth.length = depth + 2
  }

  return root
}

// --- AST -> concrete syntax (parse) tree following grammar rules 1-17 ---

function unquote(s: string): string {
  if (s.length >= 2 && s.startsWith('"') && s.endsWith('"')) {
    return s.slice(1, -1)
  }
  return s
}

function getAttr(node: AstGraphNode | undefined, name: string): string {
  if (!node) return ''
  const a = node.attrs.find(x => x.label === name)
  return a ? unquote(a.value) : ''
}

function getChild(node: AstGraphNode | undefined, label: string): AstGraphNode | undefined {
  return node?.children.find(c => c.edgeLabel === label)
}

function getItems(node: AstGraphNode | undefined): AstGraphNode[] {
  return node ? node.children.filter(c => !c.edgeLabel) : []
}

let parseNodeCounter = 0
function makeNode(label: string, kind: ParseNodeKind, children: ParseNode[] = []): ParseNode {
  return { id: `pn-${++parseNodeCounter}`, label, kind, children }
}
const nt = (name: string, children: ParseNode[]) => makeNode(`<${name}>`, 'nt', children)
const tt = (text: string) => makeNode(text, 't', [])
const eps = () => makeNode('ε', 'eps', [])

interface AddPart {
  term: AstGraphNode | undefined
  op: string
}

function flattenLeftAssoc(node: AstGraphNode | undefined, ops: string[]): AddPart[] {
  if (!node) return [{ term: undefined, op: '' }]
  if (node.type === 'BinaryOpNode') {
    const op = getAttr(node, 'op')
    if (ops.includes(op)) {
      const left = getChild(node, 'left')
      const right = getChild(node, 'right')
      const list = flattenLeftAssoc(left, ops)
      list.push({ term: right, op })
      return list
    }
  }
  return [{ term: node, op: '' }]
}

function buildFactor(node: AstGraphNode | undefined): ParseNode {
  if (!node) return nt('FACTOR', [eps()])
  switch (node.type) {
    case 'IdentifierNode':
      return nt('FACTOR', [tt(getAttr(node, 'name') || '?')])
    case 'IntLiteralNode':
    case 'FloatLiteralNode':
      return nt('FACTOR', [tt(getAttr(node, 'value') || '?')])
    case 'BinaryOpNode':
      return nt('FACTOR', [tt('('), buildExpr(node), tt(')')])
    default:
      return nt('FACTOR', [eps()])
  }
}

function buildTermTail(parts: AddPart[], i: number): ParseNode {
  if (i >= parts.length) return nt('TERM_TAIL', [eps()])
  return nt('TERM_TAIL', [
    tt(parts[i].op),
    buildFactor(parts[i].term),
    buildTermTail(parts, i + 1),
  ])
}

function buildTerm(node: AstGraphNode | undefined): ParseNode {
  const parts = flattenLeftAssoc(node, ['*', '/'])
  return nt('TERM', [
    buildFactor(parts[0].term),
    buildTermTail(parts, 1),
  ])
}

function buildExprTail(parts: AddPart[], i: number): ParseNode {
  if (i >= parts.length) return nt('EXPR_TAIL', [eps()])
  return nt('EXPR_TAIL', [
    tt(parts[i].op),
    buildTerm(parts[i].term),
    buildExprTail(parts, i + 1),
  ])
}

function buildExpr(node: AstGraphNode | undefined): ParseNode {
  const parts = flattenLeftAssoc(node, ['+', '-'])
  return nt('EXPR', [
    buildTerm(parts[0].term),
    buildExprTail(parts, 1),
  ])
}

function buildParamChain(params: AstGraphNode[], i: number): ParseNode {
  const p = params[i]
  const pname = getAttr(p, 'name') || '?'
  const ptype = getAttr(getChild(p, 'type'), 'name') || '?'
  const children: ParseNode[] = [tt(pname), tt(':'), tt(ptype)]
  if (i < params.length - 1) {
    children.push(tt(','))
    children.push(buildParamChain(params, i + 1))
  }
  return nt('PARAM', children)
}

function buildReturnStmt(body: AstGraphNode | undefined): ParseNode {
  const exprNode = getChild(body, 'expr')
  return nt('RETURN_STMT', [
    tt('return'),
    nt('RETURN_EXPR', [
      buildExpr(exprNode),
      nt('BODY_CLOSE', [
        tt('}'),
        nt('END', [tt(';')]),
      ]),
    ]),
  ])
}

function buildParseTree(ast: AstGraphNode | null): ParseNode | null {
  parseNodeCounter = 0
  if (!ast || ast.type !== 'FunctionDeclNode') return null

  const name = getAttr(ast, 'name') || '?'
  const paramsNode = getChild(ast, 'params')
  const returnTypeNode = getChild(ast, 'returnType')
  const bodyNode = getChild(ast, 'body')
  const params = getItems(paramsNode)
  const returnTypeName = getAttr(returnTypeNode, 'name') || '?'

  const paramListContents: ParseNode[] =
    params.length === 0 ? [eps()] : [buildParamChain(params, 0)]

  return nt('FUNCTION', [
    tt('fun'),
    nt('FUNC_NAME', [
      tt(name),
      nt('PARAM_LIST_START', [
        tt('('),
        nt('PARAM_LIST', [
          ...paramListContents,
          nt('PARAM_LIST_END', [
            tt(')'),
            nt('RETURN_TYPE', [
              tt(':'),
              tt(returnTypeName),
              nt('BODY_OPEN', [
                tt('{'),
                buildReturnStmt(bodyNode),
              ]),
            ]),
          ]),
        ]),
      ]),
    ]),
  ])
}

// --- Layout ---

function countNodes(node: ParseNode): number {
  return 1 + node.children.reduce((sum, c) => sum + countNodes(c), 0)
}

function buildLayout(node: ParseNode): ParseLayoutNode {
  const children = node.children.map(buildLayout)
  const width = Math.max(node.label.length * CHAR_WIDTH + 8, 28)
  const height = NODE_HEIGHT
  const childrenWidth =
    children.reduce((sum, c) => sum + c.subtreeWidth, 0) +
    Math.max(0, children.length - 1) * SIBLING_GAP
  const subtreeWidth = Math.max(width, childrenWidth)
  return {
    ...node,
    children,
    width,
    height,
    subtreeWidth,
    x: 0,
    y: 0,
  }
}

function positionLayout(node: ParseLayoutNode, left: number, depth: number) {
  node.y = DIAGRAM_PADDING_Y + depth * LEVEL_GAP
  if (node.children.length === 0) {
    node.x = left + node.subtreeWidth / 2
    return
  }
  const childrenWidth =
    node.children.reduce((sum, c) => sum + c.subtreeWidth, 0) +
    Math.max(0, node.children.length - 1) * SIBLING_GAP
  let cursor = left + (node.subtreeWidth - childrenWidth) / 2
  for (const child of node.children) {
    positionLayout(child, cursor, depth + 1)
    cursor += child.subtreeWidth + SIBLING_GAP
  }
  const first = node.children[0]
  const last = node.children[node.children.length - 1]
  node.x = (first.x + last.x) / 2
}

function collectNodes(node: ParseLayoutNode, acc: ParseLayoutNode[] = []): ParseLayoutNode[] {
  acc.push(node)
  for (const c of node.children) collectNodes(c, acc)
  return acc
}

function collectEdges(node: ParseLayoutNode, acc: ParseEdge[] = []): ParseEdge[] {
  for (const child of node.children) {
    acc.push({
      id: `${node.id}-${child.id}`,
      fromX: node.x,
      fromY: node.y + node.height - 4,
      toX: child.x,
      toY: child.y - 2,
    })
    collectEdges(child, acc)
  }
  return acc
}

function measureHeight(node: ParseLayoutNode): number {
  return Math.max(
    node.y + node.height,
    ...node.children.map(measureHeight),
  )
}

function createLayout(root: ParseNode | null): ParseLayoutResult | null {
  if (!root) return null
  const layoutRoot = buildLayout(root)
  positionLayout(layoutRoot, DIAGRAM_PADDING_X, 0)
  const nodes = collectNodes(layoutRoot)
  const edges = collectEdges(layoutRoot)
  return {
    width: layoutRoot.subtreeWidth + DIAGRAM_PADDING_X * 2,
    height: measureHeight(layoutRoot) + DIAGRAM_PADDING_Y,
    nodes,
    edges,
  }
}

export function AstViewerModal({ astText, onClose }: AstViewerModalProps) {
  const { t } = useTranslation()
  const ast = useMemo(() => parseAst(astText), [astText])
  const parseTree = useMemo(() => buildParseTree(ast), [ast])
  const totalNodes = useMemo(() => (parseTree ? countNodes(parseTree) : 0), [parseTree])
  const diagram = useMemo(() => createLayout(parseTree), [parseTree])

  return (
    <Modal title={t('astViewer.title')} onClose={onClose} width={1120} height={720}>
      <div className="ast-viewer">
        {parseTree && diagram ? (
          <>
            <div className="ast-viewer__summary">
              <p className="ast-viewer__hint">{t('astViewer.hint')}</p>
              <span className="ast-viewer__count">{t('astViewer.nodeCount', { count: String(totalNodes) })}</span>
            </div>
            <div className="ast-viewer__canvas">
              <svg
                className="ast-diagram"
                width={diagram.width}
                height={diagram.height}
                viewBox={`0 0 ${diagram.width} ${diagram.height}`}
                role="img"
                aria-label={t('astViewer.title')}
              >
                {diagram.edges.map(edge => (
                  <line
                    key={edge.id}
                    className="ast-diagram__edge"
                    x1={edge.fromX}
                    y1={edge.fromY}
                    x2={edge.toX}
                    y2={edge.toY}
                  />
                ))}

                {diagram.nodes.map(node => (
                  <text
                    key={node.id}
                    className={`ast-diagram__label ast-diagram__label--${node.kind}`}
                    x={node.x}
                    y={node.y}
                    textAnchor="middle"
                    dominantBaseline="hanging"
                  >
                    {node.label}
                  </text>
                ))}
              </svg>
            </div>
          </>
        ) : (
          <div className="ast-viewer__empty">
            <strong className="ast-viewer__empty-title">{t('astViewer.empty')}</strong>
            <p className="ast-viewer__empty-text">{t('astViewer.emptyHint')}</p>
          </div>
        )}
      </div>
      <div className="modal__footer">
        <button className="btn btn--primary" onClick={onClose}>{t('modal.close')}</button>
      </div>
    </Modal>
  )
}
