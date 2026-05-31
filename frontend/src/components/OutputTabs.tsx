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

function getParam(params: Record<string, string> | undefined, key: string): string {
  return params?.[key] ?? ''
}

function CodeBlock({ value }: { value: string }) {
  return <pre className="lab7-code-block">{value || '—'}</pre>
}

function OptimizationSection({
  title,
  description,
  flow,
  status,
  changed,
  note,
  inputIR,
  outputIR,
}: {
  title: string
  description: string
  flow: string
  status: string
  changed: boolean
  note: string
  inputIR: string
  outputIR: string
}) {
  return (
    <section className="lab7-section">
      <div className="lab7-section__heading">
        <h3>{title}</h3>
        <span className={`lab7-badge ${changed ? 'lab7-badge--changed' : ''}`}>{status}</span>
      </div>
      <p>{description}</p>

      <div className="lab7-flow">
        <span>{flow}</span>
      </div>

      <div className="lab7-compare">
        <div>
          <h4>Входной IR</h4>
          <CodeBlock value={inputIR} />
        </div>
        <div>
          <h4>Выходной IR</h4>
          <CodeBlock value={outputIR} />
        </div>
      </div>

      <p className="lab7-note">{note}</p>
    </section>
  )
}

function Lab7Result({ params }: { params: Record<string, string> | undefined }) {
  const source = getParam(params, 'source')
  const ast = getParam(params, 'ast')
  const inputIR = getParam(params, 'inputIR')
  const foldOutputIR = getParam(params, 'foldOutputIR')
  const neutralOutputIR = getParam(params, 'neutralOutputIR')
  const optimizedOutputIR = getParam(params, 'optimizedOutputIR')

  return (
    <div className="lab7-result">
      <div className="lab7-result__header">
        <div>
          <h2>Дополнительное задание ЛР7</h2>
          <p>Функции: объявление и создание. Построение AST, TAC и локальные оптимизации.</p>
        </div>
      </div>

      <section className="lab7-section">
        <h3>Исходная конструкция</h3>
        <CodeBlock value={source} />
      </section>

      <section className="lab7-section">
        <h3>AST</h3>
        <CodeBlock value={ast} />
      </section>

      <section className="lab7-section">
        <h3>Входной IR в виде TAC</h3>
        <CodeBlock value={inputIR} />
      </section>

      <OptimizationSection
        title="Оптимизация 1. Свёртка констант"
        description="Если оба операнда операции являются целочисленными константами, операция вычисляется на этапе компиляции и заменяется одним литералом."
        flow="обход выражения снизу вверх → проверка двух константных операндов → вычисление результата → замена подвыражения"
        status={getParam(params, 'foldStatus')}
        changed={getParam(params, 'foldChanged') === 'true'}
        note={getParam(params, 'foldNote')}
        inputIR={inputIR}
        outputIR={foldOutputIR}
      />

      <OptimizationSection
        title="Оптимизация 2. Удаление нейтральных операций"
        description="Нейтральные операции вида x + 0, x - 0, x * 1, 1 * x и x / 1 удаляются, а умножение на 0 заменяется литералом 0."
        flow="обход выражения снизу вверх → поиск нейтрального операнда → выбор более простой формы → замена подвыражения"
        status={getParam(params, 'neutralStatus')}
        changed={getParam(params, 'neutralChanged') === 'true'}
        note={getParam(params, 'neutralNote')}
        inputIR={inputIR}
        outputIR={neutralOutputIR}
      />

      <section className="lab7-section lab7-section--final">
        <div className="lab7-section__heading">
          <h3>Итоговый IR после обеих оптимизаций</h3>
          <span className={`lab7-badge ${getParam(params, 'optimizedChanged') === 'true' ? 'lab7-badge--changed' : ''}`}>
            {getParam(params, 'optimizedStatus')}
          </span>
        </div>
        <CodeBlock value={optimizedOutputIR} />
      </section>
    </div>
  )
}

export function OutputTabs({ output, outputKey, outputParams, errors, onNavigate }: OutputTabsProps) {
  const [active, setActive] = useState<'output' | 'errors'>('output')
  const { t } = useTranslation()

  const displayOutput = outputKey ? t(outputKey, outputParams) : output
  const isLab7Result = outputKey === 'lab7.success'

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
          isLab7Result ? (
            <Lab7Result params={outputParams} />
          ) : (
            <textarea
              className="code-editor__textarea code-editor__textarea--output"
              value={displayOutput}
              readOnly
              placeholder={t('output.placeholder')}
              spellCheck={false}
            />
          )
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
