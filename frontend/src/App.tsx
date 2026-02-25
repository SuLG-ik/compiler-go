import { useState, useRef, useEffect } from 'react'
import { Menubar } from './components/Menubar'
import { Toolbar } from './components/Toolbar'
import { CodeEditor } from './components/CodeEditor'
import { StatusBar } from './components/StatusBar'
import { HelpModal } from './components/HelpModal'
import { AboutModal } from './components/AboutModal'
import { InfoModal } from './components/InfoModal'
import { UnsavedDialog } from './components/UnsavedDialog'
import {
  OpenFile,
  SaveFile,
  SaveFileAs,
  RunAnalyzer,
  GetVersion,
} from '../wailsjs/go/main/App'
import './App.css'

export type ModalType = 'help' | 'about' | 'task' | 'grammar' | 'class' | 'method' | 'testex' | 'refs' | 'srccode' | null

function basename(path: string): string {
  return path.split('/').pop() || path.split('\\').pop() || path
}

export default function App() {
  const [code, setCode] = useState('')
  const [output, setOutput] = useState('')
  const [currentFile, setCurrentFile] = useState('')
  const [dirty, setDirty] = useState(false)
  const [status, setStatus] = useState('Готово')
  const [cursorPos, setCursorPos] = useState({ row: 1, col: 1 })
  const [modal, setModal] = useState<ModalType>(null)
  const [version, setVersion] = useState('dev')
  const [unsaved, setUnsaved] = useState<{ onSave: () => void; onDiscard: () => void } | null>(null)
  const editorRef = useRef<HTMLTextAreaElement>(null)

  useEffect(() => {
    GetVersion().then(setVersion).catch(() => {})
  }, [])

  const title = (() => {
    const name = currentFile ? basename(currentFile) : 'Без названия'
    return `${dirty ? '* ' : ''}${name} — Compiler`
  })()

  useEffect(() => { document.title = title }, [title])

  function confirmUnsaved(then: () => void) {
    if (!dirty) { then(); return }
    setUnsaved({
      onSave: () => {
        setUnsaved(null)
        handleSave(then)
      },
      onDiscard: () => {
        setUnsaved(null)
        then()
      },
    })
  }

  function handleNew() {
    confirmUnsaved(() => {
      setCode('')
      setOutput('')
      setCurrentFile('')
      setDirty(false)
      setStatus('Новый документ создан')
    })
  }

  function handleOpen() {
    confirmUnsaved(() => {
      OpenFile().then(result => {
        if (!result) return
        setCode(result.content)
        setCurrentFile(result.path)
        setDirty(false)
        setOutput('')
        setStatus(`Открыт: ${basename(result.path)}`)
      }).catch(err => {
        setStatus(`Ошибка открытия: ${err}`)
      })
    })
  }

  function handleSave(done?: () => void) {
    if (!currentFile) {
      handleSaveAs(done)
      return
    }
    SaveFile(currentFile, code).then(() => {
      setDirty(false)
      setStatus(`Сохранён: ${basename(currentFile)}`)
      done?.()
    }).catch(err => setStatus(`Ошибка сохранения: ${err}`))
  }

  function handleSaveAs(done?: () => void) {
    SaveFileAs(currentFile, code).then(path => {
      if (!path) return
      setCurrentFile(path)
      setDirty(false)
      setStatus(`Сохранён: ${basename(path)}`)
      done?.()
    }).catch(err => setStatus(`Ошибка сохранения: ${err}`))
  }

  function handleExit() {
    confirmUnsaved(() => {
      const rt = (window as any)['runtime']
      if (rt?.Quit) rt.Quit()
    })
  }

  function handleUndo() { editorRef.current?.focus(); document.execCommand('undo') }
  function handleRedo() { editorRef.current?.focus(); document.execCommand('redo') }
  function handleCut() { editorRef.current?.focus(); document.execCommand('cut') }
  function handleCopy() { editorRef.current?.focus(); document.execCommand('copy') }
  function handlePaste() {
    editorRef.current?.focus()
    navigator.clipboard.readText().then(text => {
      const el = editorRef.current
      if (!el) return
      const start = el.selectionStart
      const end = el.selectionEnd
      const newVal = code.slice(0, start) + text + code.slice(end)
      setCode(newVal)
      setDirty(true)
      setTimeout(() => {
        el.selectionStart = el.selectionEnd = start + text.length
      }, 0)
    })
  }
  function handleDelete() {
    const el = editorRef.current
    if (!el) return
    el.focus()
    const start = el.selectionStart
    const end = el.selectionEnd
    if (start !== end) {
      const newVal = code.slice(0, start) + code.slice(end)
      setCode(newVal)
      setDirty(true)
    }
  }
  function handleSelectAll() {
    const el = editorRef.current
    if (!el) return
    el.focus()
    el.select()
  }

  function handleRun() {
    setStatus('Выполняется анализ...')
    RunAnalyzer(code).then(result => {
      setOutput(result)
      setStatus('Анализ завершён')
    }).catch(err => {
      setOutput(`Ошибка: ${err}`)
      setStatus('Ошибка анализа')
    })
  }

  function handleTask() { setModal('task') }
  function handleGrammar() { setModal('grammar') }
  function handleClass() { setModal('class') }
  function handleMethod() { setModal('method') }
  function handleTestEx() { setModal('testex') }
  function handleRefs() { setModal('refs') }
  function handleSrcCode() { setModal('srccode') }

  function handleHelp() { setModal('help') }
  function handleAbout() { setModal('about') }

  function handleCodeChange(val: string) {
    setCode(val)
    if (!dirty) setDirty(true)
    setStatus('Изменён')
  }

  return (
    <div className="app-shell">
      <Menubar
        onNew={handleNew}
        onOpen={handleOpen}
        onSave={() => handleSave()}
        onSaveAs={() => handleSaveAs()}
        onExit={handleExit}
        onUndo={handleUndo}
        onRedo={handleRedo}
        onCut={handleCut}
        onCopy={handleCopy}
        onPaste={handlePaste}
        onDelete={handleDelete}
        onSelectAll={handleSelectAll}
        onTask={handleTask}
        onGrammar={handleGrammar}
        onClass={handleClass}
        onMethod={handleMethod}
        onTestEx={handleTestEx}
        onRefs={handleRefs}
        onSrcCode={handleSrcCode}
        onRun={handleRun}
        onHelp={handleHelp}
        onAbout={handleAbout}
      />
      <Toolbar
        onNew={handleNew}
        onOpen={handleOpen}
        onSave={() => handleSave()}
        onUndo={handleUndo}
        onRedo={handleRedo}
        onCopy={handleCopy}
        onCut={handleCut}
        onPaste={handlePaste}
        onRun={handleRun}
        onHelp={handleHelp}
        onAbout={handleAbout}
      />
      <CodeEditor
        code={code}
        output={output}
        editorRef={editorRef}
        onChange={handleCodeChange}
        onCursorChange={setCursorPos}
      />
      <StatusBar status={status} row={cursorPos.row} col={cursorPos.col} />

      {unsaved && (
        <UnsavedDialog
          onSave={unsaved.onSave}
          onDiscard={unsaved.onDiscard}
          onCancel={() => setUnsaved(null)}
        />
      )}

      {modal === 'help' && <HelpModal onClose={() => setModal(null)} />}
      {modal === 'about' && <AboutModal version={version} onClose={() => setModal(null)} />}
      {(modal === 'task' || modal === 'grammar' || modal === 'class' ||
        modal === 'method' || modal === 'testex' || modal === 'refs' ||
        modal === 'srccode') && (
        <InfoModal
          type={modal}
          onClose={() => setModal(null)}
        />
      )}
    </div>
  )
}
