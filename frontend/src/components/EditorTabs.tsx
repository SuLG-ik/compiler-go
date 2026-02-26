import { useTranslation } from '../i18n/I18nContext'
import './EditorTabs.css'

export interface TabInfo {
  id: number
  filename: string
  dirty: boolean
}

interface EditorTabsProps {
  tabs: TabInfo[]
  activeId: number
  onSelect: (id: number) => void
  onClose: (id: number) => void
}

export function EditorTabs({ tabs, activeId, onSelect, onClose }: EditorTabsProps) {
  const { t } = useTranslation()
  return (
    <div className="editor-tabs">
      {tabs.map(tab => (
        <div
          key={tab.id}
          className={`editor-tab ${tab.id === activeId ? 'editor-tab--active' : ''}`}
          onClick={() => onSelect(tab.id)}
        >
          {tab.dirty && <span className="editor-tab__dirty" />}
          <span className="editor-tab__label">{tab.filename}</span>
          <button
            className="editor-tab__close"
            title={t('editor.closeTab')}
            onClick={e => { e.stopPropagation(); onClose(tab.id) }}
          >
            ✕
          </button>
        </div>
      ))}
    </div>
  )
}
