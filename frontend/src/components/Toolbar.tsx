import './Toolbar.css'

interface ToolbarProps {
  onNew: () => void
  onOpen: () => void
  onSave: () => void
  onUndo: () => void
  onRedo: () => void
  onCopy: () => void
  onCut: () => void
  onPaste: () => void
  onRun: () => void
  onHelp: () => void
  onAbout: () => void
}

interface ToolBtn {
  title: string
  icon: string
  action: () => void
  separator?: boolean
}

export function Toolbar(props: ToolbarProps) {
  const buttons: ToolBtn[] = [
    { title: 'Создать (Ctrl+N)', icon: '📄', action: props.onNew },
    { title: 'Открыть (Ctrl+O)', icon: '📂', action: props.onOpen },
    { title: 'Сохранить (Ctrl+S)', icon: '💾', action: props.onSave },
    { separator: true, title: '', icon: '', action: () => {} },
    { title: 'Отменить (Ctrl+Z)', icon: '↩', action: props.onUndo },
    { title: 'Повторить (Ctrl+Y)', icon: '↪', action: props.onRedo },
    { title: 'Копировать (Ctrl+C)', icon: '⧉', action: props.onCopy },
    { title: 'Вырезать (Ctrl+X)', icon: '✂', action: props.onCut },
    { title: 'Вставить (Ctrl+V)', icon: '📋', action: props.onPaste },
    { separator: true, title: '', icon: '', action: () => {} },
    { title: 'Пуск (Ctrl+R)', icon: '▶', action: props.onRun },
    { separator: true, title: '', icon: '', action: () => {} },
    { title: 'Справка (F1)', icon: '❓', action: props.onHelp },
    { title: 'О программе', icon: 'ℹ', action: props.onAbout },
  ]

  return (
    <div className="toolbar">
      {buttons.map((btn, i) =>
        btn.separator ? (
          <div key={i} className="toolbar__sep" />
        ) : (
          <button
            key={i}
            className="toolbar__btn"
            title={btn.title}
            onClick={btn.action}
          >
            {btn.icon}
          </button>
        )
      )}
    </div>
  )
}
