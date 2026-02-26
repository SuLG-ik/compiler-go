import { Modal } from './Modal'
import { useTranslation } from '../i18n/I18nContext'
import type { ModalType } from '../App'

interface InfoModalProps {
  type: Exclude<ModalType, 'help' | 'about' | null>
  onClose: () => void
}

export function InfoModal({ type, onClose }: InfoModalProps) {
  const { t } = useTranslation()
  const title = t('info.' + type + '.title')

  return (
    <Modal title={title} onClose={onClose} width={540}>
      <p style={{ whiteSpace: 'pre-wrap', lineHeight: 1.65 }}>TODO</p>
      <div className="modal__footer" style={{ padding: '16px 0 0', border: 'none' }}>
        <button className="btn btn--primary" onClick={onClose}>{t('modal.close')}</button>
      </div>
    </Modal>
  )
}

