import { HighlightStyle, LanguageSupport, StreamLanguage, syntaxHighlighting } from '@codemirror/language'
import { tags } from '@lezer/highlight'

const KEYWORDS = new Set<string>(['fun', 'return'])
const TYPES = new Set<string>(['Int', 'Boolean', 'Float', 'Double'])

interface ParserState {
  expectFunctionName: boolean
}

const kotlinSubsetHighlight = HighlightStyle.define([
  { tag: tags.keyword, color: '#005fb8', fontWeight: '600' },
  { tag: tags.typeName, color: '#267f99', fontWeight: '600' },
  { tag: tags.definition(tags.variableName), color: '#795e26', fontWeight: '700' },
  { tag: tags.variableName, color: 'var(--text-primary)' },
  { tag: tags.number, color: '#098658' },
  { tag: tags.operator, color: '#b00020' },
  { tag: [tags.punctuation, tags.bracket], color: '#444444' },
  { tag: tags.invalid, color: '#d32f2f', textDecoration: 'underline wavy' },
])

const customLanguage = StreamLanguage.define<ParserState>({
  name: 'kotlin-subset',

  startState(): ParserState {
    return { expectFunctionName: false }
  },

  token(stream, state): string | null {
    if (stream.eatSpace()) return null

    if (stream.match(/^[0-9]+\.[0-9]+/)) {
      state.expectFunctionName = false
      return 'number'
    }

    if (stream.match(/^[0-9]+/)) {
      state.expectFunctionName = false
      return 'number'
    }

    if (stream.match(/^[A-Za-z_][A-Za-z0-9_]*/)) {
      const word = stream.current()

      if (KEYWORDS.has(word)) {
        state.expectFunctionName = word === 'fun'
        return 'keyword'
      }

      if (TYPES.has(word)) {
        state.expectFunctionName = false
        return 'type'
      }

      if (state.expectFunctionName) {
        state.expectFunctionName = false
        return 'def'
      }

      return 'variableName'
    }

    if (stream.match(/^[+\-*/]/)) {
      state.expectFunctionName = false
      return 'operator'
    }

    if (stream.match(/^[,:;]/)) {
      state.expectFunctionName = false
      return 'punctuation'
    }

    if (stream.match(/^[(){}]/)) {
      state.expectFunctionName = false
      return 'bracket'
    }

    stream.next()
    state.expectFunctionName = false
    return 'invalid'
  },
})

export function customLang(): LanguageSupport {
  return new LanguageSupport(customLanguage, [syntaxHighlighting(kotlinSubsetHighlight)])
}
