import { MutableRefObject, useCallback, useEffect, useRef, useState } from 'react'
import { EditorView, keymap, lineNumbers, drawSelection, highlightActiveLine, placeholder } from '@codemirror/view'
import { EditorState, Extension, Compartment } from '@codemirror/state'
import { defaultKeymap, historyKeymap, history } from '@codemirror/commands'
import { syntaxHighlighting, defaultHighlightStyle } from '@codemirror/language'
import { customLang } from '../lang/customLang'
import { useTranslation } from '../i18n/I18nContext'
import { OutputTabs } from './OutputTabs'
import { docStore } from '../stores/docStore'
import type { AnalyzerError } from '../hooks/useEditorState'
import './CodeEditor.css'

interface CodeEditorProps {
  tabId: number
  tabRevision: number
  output: string
  outputKey: string
  errors: AnalyzerError[]
  fontSize: number
  editorRef: MutableRefObject<EditorView | null>
  onDirty: () => void
  onCursorChange: (pos: { row: number; col: number }) => void
}

export function CodeEditor({ tabId, tabRevision, output, outputKey, errors, fontSize, editorRef, onDirty, onCursorChange }: CodeEditorProps) {
  const [splitRatio, setSplitRatio] = useState(0.72)
  const containerRef = useRef<HTMLDivElement>(null)
  const mountRef = useRef<HTMLDivElement>(null)
  const dragging = useRef(false)
  const { t } = useTranslation()
  const placeholderCompartment = useRef(new Compartment())
  const onDirtyRef = useRef(onDirty)
  const onCursorChangeRef = useRef(onCursorChange)
  useEffect(() => { onDirtyRef.current = onDirty }, [onDirty])
  useEffect(() => { onCursorChangeRef.current = onCursorChange }, [onCursorChange])
  const prevTabIdRef = useRef<number>(tabId)
  const extensionsRef = useRef<Extension[] | null>(null)
  const suppressRef = useRef(false)

  function buildExtensions() {
    return [
      history(),
      lineNumbers(),
      drawSelection(),
      highlightActiveLine(),
      syntaxHighlighting(defaultHighlightStyle, { fallback: true }),
      customLang(),
      keymap.of([...defaultKeymap, ...historyKeymap]),
      placeholderCompartment.current.of(placeholder(t('editor.placeholder'))),
      EditorView.updateListener.of(update => {
        if (update.docChanged && !suppressRef.current) {
          onDirtyRef.current()
        }
        if ((update.docChanged || update.selectionSet) && !suppressRef.current) {
          const pos = update.state.selection.main.head
          const line = update.state.doc.lineAt(pos)
          onCursorChangeRef.current({ row: line.number, col: pos - line.from + 1 })
        }
      }),
    ]
  }

  useEffect(() => {
    if (!mountRef.current) return

    const extensions = [
      ...buildExtensions(),
      EditorView.theme({
        '&': { height: '100%' },
        '.cm-scroller': {
          fontFamily: 'var(--font-mono)',
          fontSize: 'var(--editor-font-size, 14px)',
          lineHeight: '1.6',
          overflow: 'auto',
        },
        '.cm-content': { padding: '10px 4px' },
        '.cm-gutters': {
          background: 'var(--bg-toolbar)',
          color: 'var(--text-secondary)',
          border: 'none',
          borderRight: '1px solid var(--border)',
        },
        '.cm-activeLineGutter': { background: 'transparent' },
        '.cm-activeLine': { background: 'rgba(0,122,204,0.06)' },
        '.cm-cursor': { borderLeftColor: 'var(--text-primary)' },
        '.cm-selectionBackground, &.cm-focused .cm-selectionBackground': {
          background: 'rgba(0,122,204,0.25) !important',
        },
      }),
    ]
    extensionsRef.current = extensions

    const initialContent = docStore.getPending(tabId) ?? ''

    const view = new EditorView({
      state: EditorState.create({
        doc: initialContent,
        extensions,
      }),
      parent: mountRef.current,
    })

    editorRef.current = view
    docStore.save(tabId, view.state)

    return () => {
      if (editorRef.current) {
        docStore.save(prevTabIdRef.current, editorRef.current.state)
      }
      view.destroy()
      editorRef.current = null
    }
  }, [])

  useEffect(() => {
    const view = editorRef.current
    if (!view) return
    view.dispatch({
      effects: placeholderCompartment.current.reconfigure(placeholder(t('editor.placeholder'))),
    })
  }, [t])

  useEffect(() => {
    const view = editorRef.current
    if (!view) return

    const prevId = prevTabIdRef.current

    if (prevId !== tabId) {
      docStore.save(prevId, view.state)

      suppressRef.current = true
      const saved = docStore.restore(tabId)
      if (saved) {
        view.setState(saved)
      } else {
        const content = docStore.getPending(tabId) ?? ''
        view.setState(EditorState.create({
          doc: content,
          extensions: extensionsRef.current ?? [],
        }))
      }
      docStore.save(tabId, view.state)
      suppressRef.current = false
      prevTabIdRef.current = tabId

      const pos = view.state.selection.main.head
      const line = view.state.doc.lineAt(pos)
      onCursorChangeRef.current({ row: line.number, col: pos - line.from + 1 })
    } else {
      const pending = docStore.getPending(tabId)
      if (pending !== undefined) {
        suppressRef.current = true
        view.setState(EditorState.create({
          doc: pending,
          extensions: extensionsRef.current ?? [],
        }))
        docStore.save(tabId, view.state)
        suppressRef.current = false
      }
    }
  }, [tabId, tabRevision])

  const handleMouseDown = useCallback((e: React.MouseEvent) => {
    e.preventDefault()
    dragging.current = true

    function onMove(ev: MouseEvent) {
      if (!dragging.current || !containerRef.current) return
      const rect = containerRef.current.getBoundingClientRect()
      const ratio = (ev.clientY - rect.top) / rect.height
      setSplitRatio(Math.max(0.2, Math.min(0.85, ratio)))
    }
    function onUp() {
      dragging.current = false
      window.removeEventListener('mousemove', onMove)
      window.removeEventListener('mouseup', onUp)
    }
    window.addEventListener('mousemove', onMove)
    window.addEventListener('mouseup', onUp)
  }, [])

  return (
    <div className="code-editor" ref={containerRef} style={{ '--editor-font-size': `${fontSize}px` } as React.CSSProperties}>
      <div className="code-editor__pane" style={{ flex: splitRatio }}>
        <div className="code-editor__label">{t('editor.label')}</div>
        <div ref={mountRef} className="code-editor__cm" />
      </div>
      <div
        className="code-editor__divider"
        onMouseDown={handleMouseDown}
        title={t('divider.title')}
      />
      <div className="code-editor__pane code-editor__pane--output" style={{ flex: 1 - splitRatio }}>
        <OutputTabs output={output} outputKey={outputKey} errors={errors} />
      </div>
    </div>
  )
}
