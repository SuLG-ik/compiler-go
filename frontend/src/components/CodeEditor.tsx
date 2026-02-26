import { MutableRefObject, useCallback, useEffect, useRef, useState } from 'react'
import { EditorView, keymap, lineNumbers, drawSelection, highlightActiveLine, placeholder } from '@codemirror/view'
import { EditorState, Extension, Compartment } from '@codemirror/state'
import { defaultKeymap, historyKeymap, history } from '@codemirror/commands'
import { syntaxHighlighting, defaultHighlightStyle } from '@codemirror/language'
import { customLang } from '../lang/customLang'
import { useTranslation } from '../i18n/I18nContext'
import { OutputTabs } from './OutputTabs'
import type { AnalyzerError } from '../hooks/useEditorState'
import './CodeEditor.css'

interface CodeEditorProps {
  code: string
  tabId: number
  output: string
  outputKey: string
  errors: AnalyzerError[]
  fontSize: number
  editorRef: MutableRefObject<EditorView | null>
  onChange: (val: string) => void
  onCursorChange: (pos: { row: number; col: number }) => void
}

export function CodeEditor({ code, tabId, output, outputKey, errors, fontSize, editorRef, onChange, onCursorChange }: CodeEditorProps) {
  const [splitRatio, setSplitRatio] = useState(0.72)
  const containerRef = useRef<HTMLDivElement>(null)
  const mountRef = useRef<HTMLDivElement>(null)
  const dragging = useRef(false)
  const { t } = useTranslation()
  const placeholderCompartment = useRef(new Compartment())
  const onChangeRef = useRef(onChange)
  const onCursorChangeRef = useRef(onCursorChange)
  useEffect(() => { onChangeRef.current = onChange }, [onChange])
  useEffect(() => { onCursorChangeRef.current = onCursorChange }, [onCursorChange])
  const tabStatesRef = useRef<Map<number, EditorState>>(new Map())
  const prevTabIdRef = useRef<number>(tabId)
  const extensionsRef = useRef<Extension[] | null>(null)

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
        if (update.docChanged) {
          onChangeRef.current(update.state.doc.toString())
        }
        if (update.docChanged || update.selectionSet) {
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

    const view = new EditorView({
      state: EditorState.create({
        doc: code,
        extensions,
      }),
      parent: mountRef.current,
    })

    editorRef.current = view
    return () => {
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
    if (view.state.doc.toString() === code) return
    view.setState(EditorState.create({
      doc: code,
      extensions: extensionsRef.current ?? [],
    }))
  }, [code])

  useEffect(() => {
    const view = editorRef.current
    if (!view) return

    const prevId = prevTabIdRef.current
    if (prevId === tabId) return

    tabStatesRef.current.set(prevId, view.state)

    const saved = tabStatesRef.current.get(tabId)
    if (saved) {
      view.setState(saved)
    } else {
      view.setState(EditorState.create({
        doc: code,
        extensions: extensionsRef.current ?? [],
      }))
    }

    prevTabIdRef.current = tabId
  }, [tabId])

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
