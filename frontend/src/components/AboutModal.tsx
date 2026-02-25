import { Modal } from './Modal'

interface AboutModalProps {
  version: string
  onClose: () => void
}

export function AboutModal({ version, onClose }: AboutModalProps) {
  return (
    <Modal title="О программе" onClose={onClose} width={420}>
      <div style={{ textAlign: 'center', padding: '8px 0' }}>
        <div style={{ fontSize: 32, marginBottom: 12 }}>⚙️</div>
        <div style={{ fontWeight: 700, fontSize: 16, marginBottom: 6 }}>Compiler</div>
        <div style={{ color: 'var(--text-secondary)', marginBottom: 12 }}>
          Учебный текстовый редактор
        </div>
        <div style={{ marginBottom: 6 }}>Версия: <strong>{version}</strong></div>
        <div style={{ color: 'var(--text-secondary)', fontSize: 12, marginBottom: 16 }}>
          Разработан в рамках курса<br />
          «Системное программирование».<br /><br />
          Язык реализации: Go<br />
          GUI-фреймворк: Wails v2 + React
        </div>
        <button className="btn btn--primary" onClick={onClose}>OK</button>
      </div>
    </Modal>
  )
}
