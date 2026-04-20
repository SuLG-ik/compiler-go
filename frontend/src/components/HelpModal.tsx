import { Modal } from './Modal'
import { useTranslation } from '../i18n/I18nContext'
import { meta } from '../commands'
import type { CommandId } from '../commands'
import './HelpModal.css'

interface HelpModalProps {
  onClose: () => void
}

type HelpRow = { id: CommandId }

const FILE_ROWS: HelpRow[] = [
  { id: 'new' }, { id: 'open' }, { id: 'save' }, { id: 'saveAs' }, { id: 'exit' },
]
const EDIT_ROWS: HelpRow[] = [
  { id: 'undo' }, { id: 'redo' }, { id: 'cut' }, { id: 'copy' },
  { id: 'paste' }, { id: 'delete' }, { id: 'selectAll' },
]
const TEXT_ROWS: HelpRow[] = [
  { id: 'task' }, { id: 'grammar' }, { id: 'class' }, { id: 'method' },
  { id: 'testex' }, { id: 'refs' }, { id: 'srccode' },
]
const RUN_ROWS: HelpRow[]  = [{ id: 'run' }, { id: 'runAntlr' }, { id: 'runSemantic' }]
const HELP_ROWS: HelpRow[] = [{ id: 'help' }, { id: 'about' }]

const TOOLBAR_KEYS = [
  'help.toolbar.file',
  'help.toolbar.edit',
  'help.toolbar.analysis',
  'help.toolbar.info',
  'help.toolbar.font',
] as const

const DOCUMENT_KEYS = [
  'help.documents.tabs',
  'help.documents.switch',
  'help.documents.unsaved',
  'help.documents.close',
] as const

const EDITOR_KEYS = [
  'help.editor.area',
  'help.editor.lines',
  'help.editor.highlight',
  'help.editor.editing',
  'help.editor.drop',
] as const

const OUTPUT_KEYS = [
  'help.output.area',
  'help.output.results',
  'help.output.errors',
  'help.output.jump',
  'help.output.divider',
] as const

const STATUS_KEYS = [
  'help.status.state',
  'help.status.cursor',
  'help.status.language',
] as const

const EXTRA_KEYS = [
  'help.extra.shortcuts',
  'help.extra.unsaved',
  'help.extra.analysis',
  'help.extra.font',
] as const

export function HelpModal({ onClose }: HelpModalProps) {
  const { t } = useTranslation()

  function renderTable(rows: HelpRow[]) {
    return (
      <table>
        <thead>
          <tr>
            <th>{t('help.col.cmd')}</th>
            <th>{t('help.col.shortcut')}</th>
            <th>{t('help.col.desc')}</th>
          </tr>
        </thead>
        <tbody>
          {rows.map(({ id }) => (
            <tr key={id}>
              <td>{t('cmd.' + id)}</td>
              <td>{meta(id).shortcut ?? '—'}</td>
              <td>{t('help.desc.' + id)}</td>
            </tr>
          ))}
        </tbody>
      </table>
    )
  }

  function renderList(keys: readonly string[]) {
    return (
      <ul>
        {keys.map(key => (
          <li key={key}>{t(key)}</li>
        ))}
      </ul>
    )
  }

  return (
    <Modal title={t('help.title')} onClose={onClose} width={820} height={620}>
      <div className="help-content">
        <h2>{t('help.heading')}</h2>
        <p className="help-content__lead">{t('help.intro')}</p>

        <h3>{t('help.section.file')}</h3>
        {renderTable(FILE_ROWS)}

        <h3>{t('help.section.edit')}</h3>
        {renderTable(EDIT_ROWS)}

        <h3>{t('help.section.text')}</h3>
        {renderTable(TEXT_ROWS)}

        <h3>{t('help.section.run')}</h3>
        {renderTable(RUN_ROWS)}

        <h3>{t('help.section.help')}</h3>
        {renderTable(HELP_ROWS)}

        <h3>{t('help.section.toolbar')}</h3>
        {renderList(TOOLBAR_KEYS)}

        <h3>{t('help.section.documents')}</h3>
        {renderList(DOCUMENT_KEYS)}

        <h3>{t('help.section.editor')}</h3>
        {renderList(EDITOR_KEYS)}

        <h3>{t('help.section.output')}</h3>
        {renderList(OUTPUT_KEYS)}

        <h3>{t('help.section.status')}</h3>
        {renderList(STATUS_KEYS)}

        <h3>{t('help.section.extra')}</h3>
        {renderList(EXTRA_KEYS)}
      </div>
      <div className="modal__footer">
        <button className="btn btn--primary" onClick={onClose}>{t('modal.close')}</button>
      </div>
    </Modal>
  )
}
