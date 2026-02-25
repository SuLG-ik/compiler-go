import { RefObject, useCallback, useRef, useState } from 'react'
import './CodeEditor.css'

interface CodeEditorProps {
  code: string
  output: string
  editorRef: RefObject<HTMLTextAreaElement>
  onChange: (val: string) => void
  onCursorChange: (pos: { row: number; col: number }) => void
}

export function CodeEditor({ code, output, editorRef, onChange, onCursorChange }: CodeEditorProps) {
  const [splitRatio, setSplitRatio] = useState(0.72)
  const containerRef = useRef<HTMLDivElement>(null)
  const dragging = useRef(false)

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

  function handleSelect() {
    const el = editorRef.current
    if (!el) return
    const text = el.value.slice(0, el.selectionStart)
    const lines = text.split('\n')
    const row = lines.length
    const col = (lines[lines.length - 1] || '').length + 1
    onCursorChange({ row, col })
  }

  function handleKeyUp() { handleSelect() }
  function handleClick() { handleSelect() }

  return (
    <div className="code-editor" ref={containerRef}>
      <div className="code-editor__pane" style={{ flex: splitRatio }}>
        <div className="code-editor__label">Редактор</div>
        <textarea
          ref={editorRef}
          className="code-editor__textarea"
          value={code}
          onChange={e => onChange(e.target.value)}
          onSelect={handleSelect}
          onKeyUp={handleKeyUp}
          onClick={handleClick}
          placeholder="Код здесь..."
          spellCheck={false}
          autoComplete="off"
          autoCorrect="off"
          autoCapitalize="off"
        />
      </div>
      <div
        className="code-editor__divider"
        onMouseDown={handleMouseDown}
        title="Перетащите для изменения размеров"
      />
      <div className="code-editor__pane code-editor__pane--output" style={{ flex: 1 - splitRatio }}>
        <div className="code-editor__label">Результаты</div>
        <textarea
          className="code-editor__textarea code-editor__textarea--output"
          value={output}
          readOnly
          placeholder="Результаты работы языкового процессора"
          spellCheck={false}
        />
      </div>
    </div>
  )
}
