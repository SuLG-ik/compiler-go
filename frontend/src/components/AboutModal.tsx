import { Modal } from './Modal'
import { useTranslation } from '../i18n/I18nContext'

interface AboutModalProps {
  version: string
  onClose: () => void
}

export function AboutModal({ version, onClose }: AboutModalProps) {
  const { t } = useTranslation()
  return (
    <Modal title={t('about.title')} onClose={onClose} width={640} height={560}>
      <div style={{ padding: '8px 0' }}>
        <div style={{ textAlign: 'center', marginBottom: 16 }}>
          <div style={{ fontSize: 32, marginBottom: 12 }}>⚙️</div>
          <div style={{ fontWeight: 700, fontSize: 16, marginBottom: 12 }}>Compiler</div>
          <div style={{ marginBottom: 6 }}>{t('about.version')}: <strong>{version}</strong></div>
        </div>
        <div style={{ color: 'var(--text-primary)', fontSize: 13, marginBottom: 16, whiteSpace: 'pre-line', lineHeight: 1.7 }}>
          {t('about.description')}
        </div>
        <div className="modal__footer" style={{ padding: '16px 0 0', border: 'none' }}>
          <button className="btn btn--primary" onClick={onClose}>{t('modal.ok')}</button>
        </div>
      </div>
    </Modal>
  )
}
