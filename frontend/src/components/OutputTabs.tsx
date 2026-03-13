import { useState } from 'react'
import { useTranslation } from '../i18n/I18nContext'
import type { AnalyzerError, Token } from '../hooks/useEditorState'
import './OutputTabs.css'

interface OutputTabsProps {
  output: string
  outputKey: string
  errors: AnalyzerError[]
  tokens: Token[]
  onNavigate?: (line: number, col: number) => void
}

function formatLexeme(lexeme: string): string {
  return lexeme
    .replace(/ /g, '␣')
    .replace(/\t/g, '→')
    .replace(/\n/g, '↵')
    .replace(/\r/g, '')
}

export function OutputTabs({ output, outputKey, errors, tokens, onNavigate }: OutputTabsProps) {
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
          {t('output.results')}{tokens.length > 0 ? ` (${tokens.length})` : ''}
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
          tokens.length === 0 ? (
            <textarea
              className="code-editor__textarea code-editor__textarea--output"
              value={displayOutput}
              readOnly
              placeholder={t('output.placeholder')}
              spellCheck={false}
            />
          ) : (
            <div className="error-table-wrap">
              <table className="error-table">
                <thead>
                  <tr>
                    <th>{t('output.colNum')}</th>
                    <th>{t('output.colCode')}</th>
                    <th>{t('output.colType')}</th>
                    <th>{t('output.colLexeme')}</th>
                    <th>{t('output.colLine')}</th>
                    <th>{t('output.colStart')}</th>
                    <th>{t('output.colEnd')}</th>
                  </tr>
                </thead>
                <tbody>
                  {tokens.map((tok, i) => (
                    <tr
                      key={i}
                      className="error-table__row--clickable"
                      onClick={() => onNavigate?.(tok.line, tok.startPos)}
                    >
                      <td>{i + 1}</td>
                      <td>{tok.code}</td>
                      <td>{t(tok.typeKey)}</td>
                      <td className="error-table__lexeme">{formatLexeme(tok.lexeme)}</td>
                      <td>{tok.line}</td>
                      <td>{tok.startPos}</td>
                      <td>{tok.endPos}</td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          )
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
                    <tr
                      key={i}
                      className="error-table__row--clickable"
                      onClick={() => onNavigate?.(err.line, err.col)}
                    >
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
