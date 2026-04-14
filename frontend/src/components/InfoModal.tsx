import { useState } from 'react'
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
  const body = t('info.' + type + '.body')
  const [copied, setCopied] = useState(false)

  async function onCopy() {
    try {
      if (navigator.clipboard?.writeText) {
        await navigator.clipboard.writeText(body)
      } else {
        const textarea = document.createElement('textarea')
        textarea.value = body
        textarea.setAttribute('readonly', '')
        textarea.style.position = 'absolute'
        textarea.style.left = '-9999px'
        document.body.appendChild(textarea)
        textarea.select()
        document.execCommand('copy')
        document.body.removeChild(textarea)
      }
      setCopied(true)
      window.setTimeout(() => setCopied(false), 1500)
    } catch {
      setCopied(false)
    }
  }

  return (
    <Modal title={title} onClose={onClose} width={640}>
      <p style={{ whiteSpace: 'pre-wrap', lineHeight: 1.65, userSelect: 'text', WebkitUserSelect: 'text' }}>{body}</p>
      <div className="modal__footer" style={{ padding: '16px 0 0', border: 'none' }}>
        <button className="btn btn--default" onClick={onCopy}>{copied ? t('info.copied') : t('info.copy')}</button>
        <button className="btn btn--primary" onClick={onClose}>{t('modal.close')}</button>
      </div>
    </Modal>
  )
}

