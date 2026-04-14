import { useState, useRef, useEffect } from 'react'
import { useTranslation } from '../i18n/I18nContext'
import './Menubar.css'
import { type CommandRegistry, cmd, meta } from '../commands'

interface MenubarProps {
  commands: CommandRegistry
}

interface MenuItem {
  id?: string
  label: string
  shortcut?: string
  action?: () => void
  separator?: boolean
}

interface MenuDef {
  label: string
  items: MenuItem[]
}

export function Menubar({ commands }: MenubarProps) {
  const [openMenu, setOpenMenu] = useState<string | null>(null)
  const ref = useRef<HTMLDivElement>(null)
  const { t } = useTranslation()

  useEffect(() => {
    function onClick(e: MouseEvent) {
      if (ref.current && !ref.current.contains(e.target as Node)) {
        setOpenMenu(null)
      }
    }
    document.addEventListener('mousedown', onClick)
    return () => document.removeEventListener('mousedown', onClick)
  }, [])

  function exec(id: Parameters<typeof cmd>[1]) {
    setOpenMenu(null)
    cmd(commands, id)()
  }

  function item(id: Parameters<typeof cmd>[1], overrideLabel?: string): MenuItem {
    const m = meta(id)
    return { id, label: overrideLabel ?? t('cmd.' + id), shortcut: m.shortcut, action: () => exec(id) }
  }

  const menus: MenuDef[] = [
    {
      label: t('menu.file'),
      items: [
        item('new'),
        item('open'),
        item('save'),
        item('saveAs'),
        { separator: true, label: '' },
        item('exit'),
      ],
    },
    {
      label: t('menu.edit'),
      items: [
        item('undo'),
        item('redo'),
        { separator: true, label: '' },
        item('cut'),
        item('copy'),
        item('paste'),
        item('delete'),
        { separator: true, label: '' },
        item('selectAll'),
      ],
    },
    {
      label: t('menu.text'),
      items: [
        item('task'),
        item('grammar'),
        item('class'),
        item('method'),
        item('testex'),
        item('refs'),
        item('srccode'),
      ],
    },
    {
      label: t('menu.run'),
      items: [item('run', t('cmd.run'))],
    },
    {
      label: t('menu.help'),
      items: [item('help'), item('about')],
    },
  ]

  useEffect(() => {
    function onKey(e: KeyboardEvent) {
      const ctrl = e.ctrlKey || e.metaKey
      if (ctrl && e.key === 'n') { e.preventDefault(); cmd(commands, 'new')() }
      if (ctrl && e.key === 'o') { e.preventDefault(); cmd(commands, 'open')() }
      if (ctrl && !e.shiftKey && e.key === 's') { e.preventDefault(); cmd(commands, 'save')() }
      if (ctrl && e.shiftKey && e.key === 'S') { e.preventDefault(); cmd(commands, 'saveAs')() }
      if (ctrl && e.key === 'r') { e.preventDefault(); cmd(commands, 'run')() }
      if (ctrl && (e.key === '=' || e.key === '+')) { e.preventDefault(); cmd(commands, 'fontSizeUp')() }
      if (ctrl && e.key === '-') { e.preventDefault(); cmd(commands, 'fontSizeDown')() }
      if (e.key === 'F1') { e.preventDefault(); cmd(commands, 'help')() }
    }
    window.addEventListener('keydown', onKey)
    return () => window.removeEventListener('keydown', onKey)
  }, [commands])

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
