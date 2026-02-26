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
const RUN_ROWS: HelpRow[]  = [{ id: 'run' }]
const HELP_ROWS: HelpRow[] = [{ id: 'help' }, { id: 'about' }]

const UI_KEYS = ['editor', 'output', 'divider', 'statusbar'] as const

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

  return (
    <Modal title={t('help.title')} onClose={onClose} width={720} height={520}>
      <div className="help-content">
        <h2>{t('help.heading')}</h2>

        <h3>{t('help.section.file')}</h3>
        {renderTable(FILE_ROWS)}

        <h3>{t('help.section.edit')}</h3>
        {renderTable(EDIT_ROWS)}

        <h3>{t('help.section.run')}</h3>
        {renderTable(RUN_ROWS)}

        <h3>{t('help.section.help')}</h3>
        {renderTable(HELP_ROWS)}

        <h3>{t('help.section.ui')}</h3>
        <ul>
          {UI_KEYS.map(key => (
            <li key={key}>{t('help.ui.' + key)}</li>
          ))}
        </ul>
      </div>
      <div className="modal__footer">
        <button className="btn btn--primary" onClick={onClose}>{t('modal.close')}</button>
      </div>
    </Modal>
  )
}
