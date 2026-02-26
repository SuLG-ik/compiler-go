import './Toolbar.css'
import { useTranslation } from '../i18n/I18nContext'
import { type CommandRegistry, type CommandId, cmd, meta } from '../commands'

interface ToolbarProps {
  commands: CommandRegistry
}

type ToolbarItem =
  | { separator: true }
  | { id: CommandId; titleSuffix?: string }

const TOOLBAR_ITEMS: ToolbarItem[] = [
  { id: 'new' },
  { id: 'open' },
  { id: 'save' },
  { separator: true },
  { id: 'undo' },
  { id: 'redo' },
  { id: 'copy' },
  { id: 'cut' },
  { id: 'paste' },
  { separator: true },
  { id: 'run' },
  { separator: true },
  { id: 'fontSizeDown' },
  { id: 'fontSizeUp' },
  { separator: true },
  { id: 'help' },
  { id: 'about' },
]

export function Toolbar({ commands }: ToolbarProps) {
  const { t } = useTranslation()
  return (
    <div className="toolbar">
      {TOOLBAR_ITEMS.map((item, i) => {
        if ('separator' in item) {
          return <div key={i} className="toolbar__sep" />
        }
        const m = meta(item.id)
        const label = t('cmd.' + item.id)
        const title = m.shortcut ? `${label} (${m.shortcut})` : label
        return (
          <button
            key={item.id}
            className="toolbar__btn"
            title={title}
            onClick={cmd(commands, item.id)}
          >
            {m.icon}
          </button>
        )
      })}
    </div>
  )
}
