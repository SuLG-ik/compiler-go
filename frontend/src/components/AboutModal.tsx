import { Modal } from './Modal'
import { useTranslation } from '../i18n/I18nContext'

interface AboutModalProps {
  version: string
  onClose: () => void
}

export function AboutModal({ version, onClose }: AboutModalProps) {
  const { t } = useTranslation()
  return (
    <Modal title={t('about.title')} onClose={onClose} width={420}>
      <div style={{ textAlign: 'center', padding: '8px 0' }}>
        <div style={{ fontSize: 32, marginBottom: 12 }}>⚙️</div>
        <div style={{ fontWeight: 700, fontSize: 16, marginBottom: 12 }}>Compiler</div>
        <div style={{ marginBottom: 6 }}>{t('about.version')}: <strong>{version}</strong></div>
        <div style={{ color: 'var(--text-secondary)', fontSize: 12, marginBottom: 16, whiteSpace: 'pre-line' }}>
          {t('about.description')}
        </div>
        <button className="btn btn--primary" onClick={onClose}>{t('modal.ok')}</button>
      </div>
    </Modal>
  )
}
