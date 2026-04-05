import { useState } from 'react'
import { useTranslation } from '../i18n/I18nContext'
import type { AnalyzerError } from '../hooks/useEditorState'
import { REGEX_SEARCH_TASKS, type RegexSearchMatch, type RegexSearchTaskId } from '../regex/regexSearch'
import './OutputTabs.css'

interface OutputTabsProps {
  output: string
  outputKey: string
  errors: AnalyzerError[]
  regexTaskId: RegexSearchTaskId
  regexMatches: RegexSearchMatch[]
  selectedRegexMatchId: number | null
  regexMessageKey: string
  onRegexTaskChange: (taskId: RegexSearchTaskId) => void
  onRunRegexSearch: (taskId: RegexSearchTaskId) => void
  onSelectRegexMatch: (match: RegexSearchMatch) => void
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

export function OutputTabs({
  output,
  outputKey,
  errors,
  regexTaskId,
  regexMatches,
  selectedRegexMatchId,
  regexMessageKey,
  onRegexTaskChange,
  onRunRegexSearch,
  onSelectRegexMatch,
  onNavigate,
}: OutputTabsProps) {
  const [active, setActive] = useState<'output' | 'errors' | 'search'>('output')
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
          {t('output.errors')}
        </button>
        <button
          className={`output-tabs__tab ${active === 'search' ? 'output-tabs__tab--active' : ''}`}
          onClick={() => setActive('search')}
        >
          {t('output.search')}
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
                      <td>{err.messageKey ? t(err.messageKey) : err.message}</td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          )
        )}

        {active === 'search' && (
          <div className="regex-search">
            <div className="regex-search__controls">
              <label className="regex-search__field">
                <span className="regex-search__label">{t('regex.selectTask')}</span>
                <select
                  className="regex-search__select"
                  value={regexTaskId}
                  onChange={e => onRegexTaskChange(e.target.value as RegexSearchTaskId)}
                >
                  {REGEX_SEARCH_TASKS.map(task => (
                    <option key={task.id} value={task.id}>
                      {t(task.titleKey)}
                    </option>
                  ))}
                </select>
              </label>

              <button
                className="regex-search__run"
                onClick={() => onRunRegexSearch(regexTaskId)}
              >
                {t('regex.run')}
              </button>
            </div>

            <div className="error-table__summary">
              {t('regex.matchCount')}: {regexMatches.length}
            </div>

            {regexMessageKey ? (
              <div className="error-table__empty">{t(regexMessageKey)}</div>
            ) : regexMatches.length === 0 ? (
              <div className="error-table__empty">{t('regex.noMatches')}</div>
            ) : (
              <div className="error-table-wrap">
                <table className="error-table">
                  <thead>
                    <tr>
                      <th>{t('output.colFragment')}</th>
                      <th>{t('regex.colStart')}</th>
                      <th>{t('regex.colLength')}</th>
                    </tr>
                  </thead>
                  <tbody>
                    {regexMatches.map(match => (
                      <tr
                        key={match.id}
                        className={`error-table__row--clickable ${selectedRegexMatchId === match.id ? 'error-table__row--selected' : ''}`}
                        onClick={() => onSelectRegexMatch(match)}
                      >
                        <td className="error-table__lexeme">{formatLexeme(match.value)}</td>
                        <td>{formatLocation(match.line, match.col)}</td>
                        <td>{match.length}</td>
                      </tr>
                    ))}
                  </tbody>
                </table>
              </div>
            )}
          </div>
        )}
      </div>
    </div>
  )
}
