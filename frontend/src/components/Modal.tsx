import { ReactNode, useEffect } from 'react'
import './Modal.css'

interface ModalProps {
  title: string
  onClose: () => void
  children: ReactNode
  width?: number
  height?: number
}

export function Modal({ title, onClose, children, width = 560, height }: ModalProps) {
  useEffect(() => {
    function onKey(e: KeyboardEvent) {
      if (e.key === 'Escape') onClose()
    }
    window.addEventListener('keydown', onKey)
    return () => window.removeEventListener('keydown', onKey)
  }, [onClose])

  return (
    <div className="modal-overlay" onClick={e => e.target === e.currentTarget && onClose()}>
      <div className="modal" style={{ width, ...(height ? { height } : {}) }}>
        <div className="modal__header">
          <span className="modal__title">{title}</span>
          <button className="modal__close" onClick={onClose} title="Закрыть">✕</button>
        </div>
        <div className="modal__body">
          {children}
        </div>
      </div>
    </div>
  )
}
