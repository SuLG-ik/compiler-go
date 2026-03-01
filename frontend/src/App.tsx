import { useState, useRef, useEffect } from 'react'
import { EditorView } from '@codemirror/view'
import { OnFileDrop, OnFileDropOff } from '../wailsjs/runtime/runtime'
import { Menubar } from './components/Menubar'
import { Toolbar } from './components/Toolbar'
import { EditorTabs } from './components/EditorTabs'
import { CodeEditor } from './components/CodeEditor'
import { StatusBar } from './components/StatusBar'
import { HelpModal } from './components/HelpModal'
import { AboutModal } from './components/AboutModal'
import { InfoModal } from './components/InfoModal'
import { UnsavedDialog } from './components/UnsavedDialog'
import { useEditorState } from './hooks/useEditorState'
import { basename } from './hooks/useEditorState'
import { useFileActions } from './hooks/useFileActions'
import { useTranslation } from './i18n/I18nContext'
import './App.css'

export type ModalType = 'help' | 'about' | 'task' | 'grammar' | 'class' | 'method' | 'testex' | 'refs' | 'srccode' | null

export default function App() {
  const editorRef = useRef<EditorView | null>(null)
  const state = useEditorState()
  const actions = useFileActions(state, editorRef)
  const { t } = useTranslation()

  const [modal, setModal] = useState<ModalType>(null)

  useEffect(() => {
    const name = state.currentFile ? basename(state.currentFile) : t('editor.untitled')
    document.title = `${state.dirty ? '* ' : ''}${name} — Compiler`
  }, [state.currentFile, state.dirty, t])

  const fileDropRef = useRef(actions.handleFileDrop)
  useEffect(() => { fileDropRef.current = actions.handleFileDrop })
  useEffect(() => {
    OnFileDrop((_x: number, _y: number, paths: string[]) => {
      if (paths.length > 0) fileDropRef.current(paths[0])
    }, true)
    return () => { OnFileDropOff() }
  }, [])

  const commands = {
    new: actions.handleNew,
    open: actions.handleOpen,
    save: () => actions.handleSave(),
    saveAs: () => actions.handleSaveAs(),
    exit: actions.handleExit,
    undo: actions.handleUndo,
    redo: actions.handleRedo,
    cut: actions.handleCut,
    copy: actions.handleCopy,
    paste: actions.handlePaste,
    delete: actions.handleDelete,
    selectAll: actions.handleSelectAll,
    task:    () => setModal('task'),
    grammar: () => setModal('grammar'),
    class:   () => setModal('class'),
    method:  () => setModal('method'),
    testex:  () => setModal('testex'),
    refs:    () => setModal('refs'),
    srccode: () => setModal('srccode'),
    run:     actions.handleRun,
    fontSizeUp: state.fontSizeUp,
    fontSizeDown: state.fontSizeDown,
    help:    () => setModal('help'),
    about:   () => setModal('about'),
  }

  return (
    <div className="app-shell">
      <Menubar commands={commands} />
      <Toolbar commands={commands} />
      <EditorTabs
        tabs={state.tabs.map(tab => ({
          id: tab.id,
          filename: tab.path ? basename(tab.path) : t('editor.untitled'),
          dirty: tab.dirty,
        }))}
        activeId={state.activeTabId}
        onSelect={state.switchTab}
        onClose={actions.handleCloseTab}
      />
      <CodeEditor
        tabId={state.activeTabId}
        tabRevision={state.tabRevision}
        output={state.output}
        outputKey={state.outputKey}
        errors={state.errors}
        fontSize={state.fontSize}
        editorRef={editorRef}
        onDirty={state.handleCodeChange}
        onCursorChange={state.setCursorPos}
      />
      <StatusBar status={state.status} row={state.cursorPos.row} col={state.cursorPos.col} />

      {actions.unsaved && (
        <UnsavedDialog
          onSave={actions.unsaved.onSave}
          onDiscard={actions.unsaved.onDiscard}
          onCancel={() => actions.setUnsaved(null)}
        />
      )}

      {modal === 'help' && <HelpModal onClose={() => setModal(null)} />}
      {modal === 'about' && <AboutModal version={state.version} onClose={() => setModal(null)} />}
      {(modal === 'task' || modal === 'grammar' || modal === 'class' ||
        modal === 'method' || modal === 'testex' || modal === 'refs' ||
        modal === 'srccode') && (
        <InfoModal type={modal} onClose={() => setModal(null)} />
      )}
    </div>
  )
}
