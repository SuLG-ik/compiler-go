import { Modal } from './Modal'
import { useTranslation } from '../i18n/I18nContext'

interface UnsavedDialogProps {
  onSave: () => void
  onDiscard: () => void
  onCancel: () => void
}

export function UnsavedDialog({ onSave, onDiscard, onCancel }: UnsavedDialogProps) {
  const { t } = useTranslation()
  return (
    <Modal title={t('unsaved.title')} onClose={onCancel} width={420}>
      <p style={{ marginBottom: 16, whiteSpace: 'pre-line' }}>
        {t('unsaved.message')}
      </p>
      <div className="modal__footer" style={{ padding: 0, border: 'none' }}>
        <button className="btn btn--primary" onClick={onSave}>{t('unsaved.save')}</button>
        <button className="btn btn--danger" onClick={onDiscard}>{t('unsaved.discard')}</button>
        <button className="btn btn--default" onClick={onCancel}>{t('unsaved.cancel')}</button>
      </div>
    </Modal>
  )
}
