import { StreamLanguage, LanguageSupport } from '@codemirror/language'

const KEYWORDS = new Set<string>([
  'if', 'else', 'while', 'for', 'return', 'var', 'func'
])

interface ParserState {
  inBlockComment: boolean
}

const customLanguage = StreamLanguage.define<ParserState>({
  name: 'custom',

  startState(): ParserState {
    return { inBlockComment: false }
  },

  token(stream, state): string | null {
    if (state.inBlockComment) {
      if (stream.match('*/')) {
        state.inBlockComment = false
      } else {
        stream.next()
      }
      return 'comment'
    }

    if (stream.eatSpace()) return null

    if (stream.match('//')) {
      stream.skipToEnd()
      return 'lineComment'
    }

    if (stream.match('/*')) {
      state.inBlockComment = true
      return 'comment'
    }

    if (stream.match('"')) {
      while (!stream.eol()) {
        const ch = stream.next()
        if (ch === '\\') { stream.next(); continue }
        if (ch === '"') break
      }
      return 'string'
    }

    if (stream.match("'")) {
      while (!stream.eol()) {
        const ch = stream.next()
        if (ch === '\\') { stream.next(); continue }
        if (ch === "'") break
      }
      return 'string'
    }

    if (stream.match(/^[0-9]+(\.[0-9]+)?([eE][+-]?[0-9]+)?/)) {
      return 'number'
    }

    if (stream.match(/^[A-Za-z_][A-Za-z0-9_]*/)) {
      const word = stream.current()
      if (KEYWORDS.has(word)) return 'keyword'
      return 'variableName'
    }

    if (stream.match(/^[+\-*/%=<>!&|^~?:;,.()\[\]{}]/)) {
      return 'operator'
    }

    stream.next()
    return null
  },
})

export function customLang(): LanguageSupport {
  return new LanguageSupport(customLanguage)
}
