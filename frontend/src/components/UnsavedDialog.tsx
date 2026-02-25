import { Modal } from './Modal'

interface UnsavedDialogProps {
  onSave: () => void
  onDiscard: () => void
  onCancel: () => void
}

export function UnsavedDialog({ onSave, onDiscard, onCancel }: UnsavedDialogProps) {
  return (
    <Modal title="Несохранённые изменения" onClose={onCancel} width={420}>
      <p style={{ marginBottom: 16 }}>
        Файл содержит несохранённые изменения.<br />
        Сохранить перед продолжением?
      </p>
      <div className="modal__footer" style={{ padding: 0, border: 'none' }}>
        <button className="btn btn--primary" onClick={onSave}>Сохранить</button>
        <button className="btn btn--danger" onClick={onDiscard}>Не сохранять</button>
        <button className="btn btn--default" onClick={onCancel}>Отмена</button>
      </div>
    </Modal>
  )
}
