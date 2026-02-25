import { useState, useRef, useEffect } from 'react'
import './Menubar.css'

interface MenubarProps {
  onNew: () => void
  onOpen: () => void
  onSave: () => void
  onSaveAs: () => void
  onExit: () => void
  onUndo: () => void
  onRedo: () => void
  onCut: () => void
  onCopy: () => void
  onPaste: () => void
  onDelete: () => void
  onSelectAll: () => void
  onTask: () => void
  onGrammar: () => void
  onClass: () => void
  onMethod: () => void
  onTestEx: () => void
  onRefs: () => void
  onSrcCode: () => void
  onRun: () => void
  onHelp: () => void
  onAbout: () => void
}

interface MenuItem {
  label: string
  shortcut?: string
  action?: () => void
  separator?: boolean
}

interface MenuDef {
  label: string
  items: MenuItem[]
}

export function Menubar(props: MenubarProps) {
  const [openMenu, setOpenMenu] = useState<string | null>(null)
  const ref = useRef<HTMLDivElement>(null)

  useEffect(() => {
    function onClick(e: MouseEvent) {
      if (ref.current && !ref.current.contains(e.target as Node)) {
        setOpenMenu(null)
      }
    }
    document.addEventListener('mousedown', onClick)
    return () => document.removeEventListener('mousedown', onClick)
  }, [])

  function exec(fn: () => void) {
    setOpenMenu(null)
    fn()
  }

  const menus: MenuDef[] = [
    {
      label: 'Файл',
      items: [
        { label: 'Создать', shortcut: 'Ctrl+N', action: () => exec(props.onNew) },
        { label: 'Открыть', shortcut: 'Ctrl+O', action: () => exec(props.onOpen) },
        { label: 'Сохранить', shortcut: 'Ctrl+S', action: () => exec(props.onSave) },
        { label: 'Сохранить как', shortcut: 'Ctrl+Shift+S', action: () => exec(props.onSaveAs) },
        { separator: true, label: '' },
        { label: 'Выход', action: () => exec(props.onExit) },
      ],
    },
    {
      label: 'Правка',
      items: [
        { label: 'Отменить', shortcut: 'Ctrl+Z', action: () => exec(props.onUndo) },
        { label: 'Повторить', shortcut: 'Ctrl+Y', action: () => exec(props.onRedo) },
        { separator: true, label: '' },
        { label: 'Вырезать', shortcut: 'Ctrl+X', action: () => exec(props.onCut) },
        { label: 'Копировать', shortcut: 'Ctrl+C', action: () => exec(props.onCopy) },
        { label: 'Вставить', shortcut: 'Ctrl+V', action: () => exec(props.onPaste) },
        { label: 'Удалить', action: () => exec(props.onDelete) },
        { separator: true, label: '' },
        { label: 'Выделить все', shortcut: 'Ctrl+A', action: () => exec(props.onSelectAll) },
      ],
    },
    {
      label: 'Текст',
      items: [
        { label: 'Постановка задачи', action: () => exec(props.onTask) },
        { label: 'Грамматика', action: () => exec(props.onGrammar) },
        { label: 'Классификация грамматики', action: () => exec(props.onClass) },
        { label: 'Метод анализа', action: () => exec(props.onMethod) },
        { label: 'Тестовый пример', action: () => exec(props.onTestEx) },
        { label: 'Список литературы', action: () => exec(props.onRefs) },
        { label: 'Исходный код программы', action: () => exec(props.onSrcCode) },
      ],
    },
    {
      label: 'Пуск',
      items: [
        { label: 'Пуск', shortcut: 'Ctrl+R', action: () => exec(props.onRun) },
      ],
    },
    {
      label: 'Справка',
      items: [
        { label: 'Вызов справки', shortcut: 'F1', action: () => exec(props.onHelp) },
        { label: 'О программе', action: () => exec(props.onAbout) },
      ],
    },
  ]

  useEffect(() => {
    function onKey(e: KeyboardEvent) {
      const ctrl = e.ctrlKey || e.metaKey
      if (ctrl && e.key === 'n') { e.preventDefault(); props.onNew() }
      if (ctrl && e.key === 'o') { e.preventDefault(); props.onOpen() }
      if (ctrl && !e.shiftKey && e.key === 's') { e.preventDefault(); props.onSave() }
      if (ctrl && e.shiftKey && e.key === 'S') { e.preventDefault(); props.onSaveAs() }
      if (ctrl && e.key === 'r') { e.preventDefault(); props.onRun() }
      if (e.key === 'F1') { e.preventDefault(); props.onHelp() }
    }
    window.addEventListener('keydown', onKey)
    return () => window.removeEventListener('keydown', onKey)
  }, [props])

  return (
    <nav className="menubar" ref={ref}>
      {menus.map(menu => (
        <div key={menu.label} className="menubar__item">
          <button
            className={`menubar__btn ${openMenu === menu.label ? 'menubar__btn--open' : ''}`}
            onClick={() => setOpenMenu(openMenu === menu.label ? null : menu.label)}
          >
            {menu.label}
          </button>
          {openMenu === menu.label && (
            <ul className="menubar__dropdown">
              {menu.items.map((item, i) =>
                item.separator ? (
                  <li key={i} className="menubar__sep" />
                ) : (
                  <li key={i}>
                    <button className="menubar__entry" onClick={item.action}>
                      <span>{item.label}</span>
                      {item.shortcut && <span className="menubar__shortcut">{item.shortcut}</span>}
                    </button>
                  </li>
                )
              )}
            </ul>
          )}
        </div>
      ))}
    </nav>
  )
}
