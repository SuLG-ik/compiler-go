import { EditorState } from '@codemirror/state'

const cmStates = new Map<number, EditorState>()
const pendingContent = new Map<number, string>()

export const docStore = {
  save(tabId: number, state: EditorState) {
    cmStates.set(tabId, state)
    pendingContent.delete(tabId)
  },

  restore(tabId: number): EditorState | undefined {
    return cmStates.get(tabId)
  },

  setPending(tabId: number, content: string) {
    pendingContent.set(tabId, content)
    cmStates.delete(tabId)
  },

  getPending(tabId: number): string | undefined {
    return pendingContent.get(tabId)
  },

  getCode(tabId: number): string {
    const s = cmStates.get(tabId)
    if (s) return s.doc.toString()
    return pendingContent.get(tabId) ?? ''
  },

  remove(tabId: number) {
    cmStates.delete(tabId)
    pendingContent.delete(tabId)
  },
}
