import { useState, useEffect } from 'react'
import { GetVersion } from '../../wailsjs/go/main/App'
import { docStore } from '../stores/docStore'

export function basename(path: string): string {
  return path.split('/').pop() || path.split('\\').pop() || path
}

export interface StatusMsg {
  key: string
  params?: Record<string, string>
}

export interface AnalyzerError {
  line: number
  col: number
  message: string
  messageKey?: string
}

export interface Tab {
  id: number
  path: string
  dirty: boolean
}

let nextTabId = 1

export interface EditorState {
  currentFile: string
  dirty: boolean
  tabs: Tab[]
  activeTabId: number
  tabRevision: number
  addTab: (initial?: Partial<Omit<Tab, 'id'>>) => number
  removeTab: (id: number) => void
  switchTab: (id: number) => void
  updateTab: (id: number, updates: Partial<Omit<Tab, 'id'>>) => void
  isTabDirty: (id: number) => boolean
  getTab: (id: number) => Tab | undefined
  loadTabContent: (tabId: number, content: string) => void
  output: string
  outputKey: string
  errors: AnalyzerError[]
  status: StatusMsg
  cursorPos: { row: number; col: number }
  version: string
  fontSize: number
  setCurrentFile: (val: string) => void
  setDirty: (val: boolean) => void
  setOutput: (val: string) => void
  setOutputKey: (val: string) => void
  setErrors: (val: AnalyzerError[]) => void
  setStatus: (val: StatusMsg) => void
  setCursorPos: (pos: { row: number; col: number }) => void
  handleCodeChange: () => void
  fontSizeUp: () => void
  fontSizeDown: () => void
}

export function useEditorState(): EditorState {
  const [tabs, setTabs] = useState<Tab[]>(() => {
    const id = nextTabId++
    return [{ id, path: '', dirty: false }]
  })
  const [activeTabId, setActiveTabId] = useState(() => tabs[0].id)
  const [tabRevision, setTabRevision] = useState(0)
  const [output, setOutput] = useState('')
  const [outputKey, setOutputKey] = useState('')
  const [errors, setErrors] = useState<AnalyzerError[]>([])
  const [status, setStatus] = useState<StatusMsg>({ key: 'status.ready' })
  const [cursorPos, setCursorPos] = useState({ row: 1, col: 1 })
  const [version, setVersion] = useState('dev')
  const [fontSize, setFontSize] = useState(14)

  function fontSizeUp() { setFontSize(s => Math.min(24, s + 1)) }
  function fontSizeDown() { setFontSize(s => Math.max(10, s - 1)) }

  useEffect(() => { GetVersion().then(setVersion).catch(() => {}) }, [])

  const activeTab = tabs.find(t => t.id === activeTabId) ?? tabs[0]
  const currentFile = activeTab.path
  const dirty = activeTab.dirty

  function setCurrentFile(val: string) {
    setTabs(prev => prev.map(t => t.id === activeTabId ? { ...t, path: val } : t))
  }
  function setDirty(val: boolean) {
    setTabs(prev => prev.map(t => t.id === activeTabId ? { ...t, dirty: val } : t))
  }

  function addTab(initial?: Partial<Omit<Tab, 'id'>>): number {
    const id = nextTabId++
    setTabs(prev => [...prev, { id, path: '', dirty: false, ...initial }])
    setActiveTabId(id)
    return id
  }

  function removeTab(id: number) {
    docStore.remove(id)
    setTabs(prev => {
      const next = prev.filter(t => t.id !== id)
      if (next.length === 0) {
        const newId = nextTabId++
        setActiveTabId(newId)
        return [{ id: newId, path: '', dirty: false }]
      }
      if (id === activeTabId) {
        const idx = prev.findIndex(t => t.id === id)
        const newActive = next[Math.min(idx, next.length - 1)]
        setActiveTabId(newActive.id)
      }
      return next
    })
  }

  function switchTab(id: number) { setActiveTabId(id) }

  function updateTab(id: number, updates: Partial<Omit<Tab, 'id'>>) {
    setTabs(prev => prev.map(t => t.id === id ? { ...t, ...updates } : t))
  }

  function isTabDirty(id: number): boolean {
    return tabs.find(t => t.id === id)?.dirty ?? false
  }

  function getTab(id: number) { return tabs.find(t => t.id === id) }

  function loadTabContent(tabId: number, content: string) {
    docStore.setPending(tabId, content)
    if (tabId === activeTabId) {
      setTabRevision(r => r + 1)
    }
  }

  function handleCodeChange() {
    setTabs(prev => prev.map(t => t.id === activeTabId ? { ...t, dirty: true } : t))
    setStatus({ key: 'status.modified' })
  }

  return {
    currentFile, dirty,
    tabs, activeTabId, tabRevision,
    addTab, removeTab, switchTab, updateTab, isTabDirty, getTab, loadTabContent,
    output, outputKey, errors, status, cursorPos, version, fontSize,
    setCurrentFile, setDirty,
    setOutput, setOutputKey, setErrors, setStatus, setCursorPos,
    handleCodeChange, fontSizeUp, fontSizeDown,
  }
}
