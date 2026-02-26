import { useState } from 'react'
import { useTranslation } from '../i18n/I18nContext'
import type { AnalyzerError } from '../hooks/useEditorState'
import './OutputTabs.css'

interface OutputTabsProps {
  output: string
  outputKey: string
  errors: AnalyzerError[]
}

export function OutputTabs({ output, outputKey, errors }: OutputTabsProps) {
  const [active, setActive] = useState<'output' | 'errors'>('output')
  const { t } = useTranslation()

  const displayOutput = outputKey ? t(outputKey) : output

  return (
    <div className="output-tabs">
      <div className="output-tabs__bar">
        <button
          className={`output-tabs__tab ${active === 'output' ? 'output-tabs__tab--active' : ''}`}
          onClick={() => setActive('output')}
        >
          {t('output.results')}
        </button>
        <button
          className={`output-tabs__tab ${active === 'errors' ? 'output-tabs__tab--active' : ''}`}
          onClick={() => setActive('errors')}
        >
          {t('output.errors')}{errors.length > 0 ? ` (${errors.length})` : ''}
        </button>
      </div>

      <div className="output-tabs__body">
        {active === 'output' && (
          <textarea
            className="code-editor__textarea code-editor__textarea--output"
            value={displayOutput}
            readOnly
            placeholder={t('output.placeholder')}
            spellCheck={false}
          />
        )}

        {active === 'errors' && (
          errors.length === 0 ? (
            <div className="error-table__empty">{t('output.noErrors')}</div>
          ) : (
            <div className="error-table-wrap">
              <table className="error-table">
                <thead>
                  <tr>
                    <th>{t('output.colNum')}</th>
                    <th>{t('output.colLine')}</th>
                    <th>{t('output.colCol')}</th>
                    <th>{t('output.colDesc')}</th>
                  </tr>
                </thead>
                <tbody>
                  {errors.map((err, i) => (
                    <tr key={i}>
                      <td>{i + 1}</td>
                      <td>{err.line}</td>
                      <td>{err.col}</td>
                      <td>{err.messageKey ? t(err.messageKey) : err.message}</td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          )
        )}
      </div>
    </div>
  )
}
