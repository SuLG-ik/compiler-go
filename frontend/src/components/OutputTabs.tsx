import { useEffect, useState } from 'react'
import { useTranslation } from '../i18n/I18nContext'
import type { AnalyzerError } from '../hooks/useEditorState'
import './OutputTabs.css'

interface OutputTabsProps {
  output: string
  outputKey: string
  outputParams?: Record<string, string>
  errors: AnalyzerError[]
  onNavigate?: (line: number, col: number) => void
}

function formatLexeme(lexeme: string): string {
  return lexeme
    .replace(/ /g, '␣')
    .replace(/\t/g, '→')
    .replace(/\n/g, '↵')
    .replace(/\r/g, '')
}

function formatLocation(line: number, col: number): string {
  return `${line}:${col}`
}

export function OutputTabs({ output, outputKey, outputParams, errors, onNavigate }: OutputTabsProps) {
  const [active, setActive] = useState<'output' | 'errors'>('output')
  const { t } = useTranslation()

  const displayOutput = outputKey ? t(outputKey, outputParams) : output

  useEffect(() => {
    if (errors.length > 0) {
      setActive('errors')
    }
  }, [errors])

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
          {t('output.errors')}
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
              <div className="error-table__summary">
                {t('output.errorCount')}: {errors.length}
              </div>
              <table className="error-table">
                <thead>
                  <tr>
                    <th>{t('output.colFragment')}</th>
                    <th>{t('output.colLocation')}</th>
                    <th>{t('output.colDesc')}</th>
                  </tr>
                </thead>
                <tbody>
                  {errors.map((err, i) => (
                    <tr
                      key={i}
                      className="error-table__row--clickable"
                      onClick={() => onNavigate?.(err.line, err.col)}
                    >
                      <td className="error-table__lexeme">{err.fragment ? formatLexeme(err.fragment) : '—'}</td>
                      <td>{formatLocation(err.line, err.col)}</td>
                      <td>{err.messageKey ? t(err.messageKey, err.messageParams) : err.message}</td>
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
