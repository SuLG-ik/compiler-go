import { MutableRefObject, useState, useRef } from 'react'
import { EditorView } from '@codemirror/view'
import { undo, redo, selectAll } from '@codemirror/commands'
import {
  OpenFile,
  ReadFileByPath,
  SaveFile,
  SaveFileAs,
  RunAnalyzer,
} from '../../wailsjs/go/main/App'
import { docStore } from '../stores/docStore'
import { basename, type EditorState } from './useEditorState'

export interface UnsavedPrompt {
  onSave: () => void
  onDiscard: () => void
}

export function useFileActions(
  state: EditorState,
  editorRef: MutableRefObject<EditorView | null>,
) {
  const [unsaved, setUnsaved] = useState<UnsavedPrompt | null>(null)
  const handleSaveRef = useRef<(done?: () => void) => void>(() => {})

  const { currentFile, dirty,
    setOutput, setOutputKey, setErrors, setCurrentFile, setDirty, setStatus,
    tabs, activeTabId, addTab, removeTab, switchTab, updateTab, getTab, loadTabContent } = state

  function getActiveCode(): string {
    return editorRef.current?.state.doc.toString() ?? ''
  }

  function handleNew() {
    addTab()
    setOutput('')
    setOutputKey('')
    setErrors([])
    setStatus({ key: 'status.newDoc' })
  }

  function handleOpen() {
    OpenFile().then(result => {
      if (!result) return

      const existing = tabs.find(t => t.path === result.path && t.path !== '')
      if (existing) {
        switchTab(existing.id)
        setStatus({ key: 'status.switched', params: { name: basename(result.path) } })
        return
      }

      if (!currentFile && !dirty && !getActiveCode()) {
        updateTab(activeTabId, { path: result.path, dirty: false })
        loadTabContent(activeTabId, result.content)
      } else {
        const id = addTab({ path: result.path, dirty: false })
        docStore.setPending(id, result.content)
      }
      setOutput('')
      setOutputKey('')
      setErrors([])
      setStatus({ key: 'status.opened', params: { name: basename(result.path) } })
    }).catch(err => setStatus({ key: 'status.errorOpen', params: { err: String(err) } }))
  }

  function handleSave(done?: () => void) {
    if (!currentFile) { handleSaveAs(done); return }
    SaveFile(currentFile, getActiveCode()).then(() => {
      setDirty(false)
      setStatus({ key: 'status.saved', params: { name: basename(currentFile) } })
      done?.()
    }).catch(err => setStatus({ key: 'status.errorSave', params: { err: String(err) } }))
  }
  handleSaveRef.current = handleSave

  function handleSaveAs(done?: () => void) {
    SaveFileAs(currentFile, getActiveCode()).then(path => {
      if (!path) return
      setCurrentFile(path)
      setDirty(false)
      setStatus({ key: 'status.saved', params: { name: basename(path) } })
      done?.()
    }).catch(err => setStatus({ key: 'status.errorSave', params: { err: String(err) } }))
  }

  function handleExit() {
    const anyDirty = tabs.some(t => t.dirty)
    if (!anyDirty) { quit(); return }
    setUnsaved({
      onSave:    () => { setUnsaved(null); handleSaveRef.current(() => quit()) },
      onDiscard: () => { setUnsaved(null); quit() },
    })
  }

  function quit() {
    const rt = (window as any)['runtime']
    if (rt?.Quit) rt.Quit()
  }

  function handleCloseTab(id: number) {
    const tab = getTab(id)
    if (!tab) return
    if (!tab.dirty) { removeTab(id); return }

    if (activeTabId !== id) switchTab(id)
    setUnsaved({
      onSave:    () => { setUnsaved(null); handleSaveRef.current(() => removeTab(id)) },
      onDiscard: () => { setUnsaved(null); removeTab(id) },
    })
  }

  function handleFileDrop(path: string) {
    ReadFileByPath(path).then(result => {
      if (!result) return

      const existing = tabs.find(t => t.path === result.path && t.path !== '')
      if (existing) {
        switchTab(existing.id)
        setStatus({ key: 'status.switched', params: { name: basename(result.path) } })
        return
      }

      if (!currentFile && !dirty && !getActiveCode()) {
        updateTab(activeTabId, { path: result.path, dirty: false })
        loadTabContent(activeTabId, result.content)
      } else {
        const id = addTab({ path: result.path, dirty: false })
        docStore.setPending(id, result.content)
      }
      setOutput('')
      setOutputKey('')
      setErrors([])
      setStatus({ key: 'status.opened', params: { name: basename(result.path) } })
    }).catch(err => setStatus({ key: 'status.errorOpen', params: { err: String(err) } }))
  }

  function handleUndo() {
    const view = editorRef.current; if (!view) return
    undo(view); view.focus()
  }
  function handleRedo() {
    const view = editorRef.current; if (!view) return
    redo(view); view.focus()
  }
  function handleCopy() {
    const view = editorRef.current; if (!view) return
    const { from, to } = view.state.selection.main
    if (from !== to) navigator.clipboard.writeText(view.state.sliceDoc(from, to))
    view.focus()
  }
  function handleCut() {
    const view = editorRef.current; if (!view) return
    const { from, to } = view.state.selection.main
    if (from !== to) {
      navigator.clipboard.writeText(view.state.sliceDoc(from, to))
      view.dispatch({ changes: { from, to, insert: '' } })
    }
    view.focus()
  }
  function handlePaste() {
    const view = editorRef.current; if (!view) return
    view.contentDOM.focus()
    document.execCommand('paste')
  }
  function handleDelete() {
    const view = editorRef.current; if (!view) return
    const { from, to } = view.state.selection.main
    if (from !== to) view.dispatch({ changes: { from, to, insert: '' } })
    view.focus()
  }
  function handleSelectAll() {
    const view = editorRef.current; if (!view) return
    selectAll(view); view.focus()
  }

  function handleRun() {
    const activeCode = getActiveCode()
    setStatus({ key: 'status.analyzing' })
    RunAnalyzer(activeCode).then(result => {
      setOutputKey(result.outputKey ?? '')
      setOutput(result.outputKey ? '' : (result.output ?? ''))
      setErrors(result.errors ?? [])
      setStatus({ key: 'status.analyzed' })
    }).catch(err => {
      setOutputKey('')
      setOutput(String(err))
      setErrors([])
      setStatus({ key: 'status.errorAnalysis' })
    })
  }

  return {
    unsaved, setUnsaved,
    handleNew, handleOpen, handleSave, handleSaveAs, handleExit,
    handleUndo, handleRedo, handleCut, handleCopy, handlePaste,
    handleDelete, handleSelectAll,
    handleCloseTab, handleFileDrop, handleRun,
  }
}

