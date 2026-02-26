import { useTranslation } from '../i18n/I18nContext'
import type { StatusMsg } from '../hooks/useEditorState'
import './StatusBar.css'

interface StatusBarProps {
  status: StatusMsg
  row: number
  col: number
}

export function StatusBar({ status, row, col }: StatusBarProps) {
  const { t, lang, setLang } = useTranslation()
  return (
    <div className="statusbar">
      <span className="statusbar__status">{t(status.key, status.params)}</span>
      <span className="statusbar__cursor">{t('status.row')} {row}  {t('status.col')} {col}</span>
      <button
        className="statusbar__lang"
        onClick={() => setLang(lang === 'ru' ? 'en' : 'ru')}
        title={t('lang.label')}
      >
        {lang === 'ru' ? 'RU' : 'EN'}
      </button>
    </div>
  )
}
