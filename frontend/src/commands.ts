export type CommandId =
  | 'new' | 'open' | 'save' | 'saveAs' | 'exit'
  | 'undo' | 'redo' | 'cut' | 'copy' | 'paste' | 'delete' | 'selectAll'
  | 'task' | 'grammar' | 'class' | 'method' | 'testex' | 'refs' | 'srccode'
  | 'run' | 'runAntlr' | 'runSemantic' | 'showAst'
  | 'fontSizeUp' | 'fontSizeDown'
  | 'help' | 'about'

export interface CommandMeta {
  id: CommandId
  label: string
  shortcut?: string
  icon?: string
}

export const COMMANDS: CommandMeta[] = [
  { id: 'new',       label: 'Создать',                shortcut: 'Ctrl+N',       icon: '📄' },
  { id: 'open',      label: 'Открыть',                shortcut: 'Ctrl+O',       icon: '📂' },
  { id: 'save',      label: 'Сохранить',              shortcut: 'Ctrl+S',       icon: '💾' },
  { id: 'saveAs',    label: 'Сохранить как',          shortcut: 'Ctrl+Shift+S' },
  { id: 'exit',      label: 'Выход' },
  { id: 'undo',      label: 'Отменить',               shortcut: 'Ctrl+Z',       icon: '↩' },
  { id: 'redo',      label: 'Повторить',              shortcut: 'Ctrl+Y',       icon: '↪' },
  { id: 'cut',       label: 'Вырезать',               shortcut: 'Ctrl+X',       icon: '✂' },
  { id: 'copy',      label: 'Копировать',             shortcut: 'Ctrl+C',       icon: '⧉' },
  { id: 'paste',     label: 'Вставить',               shortcut: 'Ctrl+V',       icon: '📋' },
  { id: 'delete',    label: 'Удалить' },
  { id: 'selectAll', label: 'Выделить все',           shortcut: 'Ctrl+A' },
  { id: 'task',      label: 'Постановка задачи' },
  { id: 'grammar',   label: 'Грамматика' },
  { id: 'class',     label: 'Классификация грамматики' },
  { id: 'method',    label: 'Метод анализа' },
  { id: 'testex',    label: 'Тестовый пример' },
  { id: 'refs',      label: 'Список литературы' },
  { id: 'srccode',   label: 'Исходный код программы' },
  { id: 'run',       label: 'Пуск',                   shortcut: 'Ctrl+R',       icon: '▶' },
  { id: 'runAntlr',  label: 'Пуск ANTLR',             shortcut: 'Ctrl+Shift+R', icon: 'Ⓐ' },
  { id: 'runSemantic', label: 'Семантический анализ', shortcut: 'Ctrl+Alt+R',   icon: 'Σ' },
  { id: 'showAst',   label: 'Показать AST',           shortcut: 'Ctrl+Alt+A' },
  { id: 'fontSizeUp',   label: 'Увеличить шрифт',       shortcut: 'Ctrl+=',       icon: 'A+' },
  { id: 'fontSizeDown', label: 'Уменьшить шрифт',       shortcut: 'Ctrl+−',       icon: 'A−' },
  { id: 'help',      label: 'Вызов справки',          shortcut: 'F1',           icon: '❓' },
  { id: 'about',     label: 'О программе',                                      icon: 'ℹ' },
]

export type CommandRegistry = Partial<Record<CommandId, () => void>>

export function cmd(registry: CommandRegistry, id: CommandId): (() => void) {
  return registry[id] ?? (() => {})
}

export function meta(id: CommandId): CommandMeta {
  return COMMANDS.find(c => c.id === id)!
}
