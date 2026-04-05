export type RegexSearchTaskId = 'password' | 'visa' | 'visaAutomaton' | 'longitude'

export interface RegexSearchMatch {
  id: number
  value: string
  line: number
  col: number
  length: number
  from: number
  to: number
}

export interface RegexSearchTaskOption {
  id: RegexSearchTaskId
  titleKey: string
}

interface RegexTaskDefinition extends RegexSearchTaskOption {
  matchesLine: (lineText: string) => boolean
}

const PASSWORD_PATTERN = /^[А-Яа-яЁё0-9!@#$%^&*()_\-=+]{8,}$/u
const VISA_PATTERN = /^4[0-9]{12}([0-9]{3})?$/
const LONGITUDE_PATTERN = /^(([0-9]{1,2}|1[0-7][0-9])°[0-5][0-9]'[0-5][0-9]"[EW]|180°00'00"[EW])$/u

function isAsciiDigit(char: string): boolean {
  const code = char.charCodeAt(0)
  return code >= 48 && code <= 57
}

// Accepting states correspond to complete Visa numbers with 13 or 16 digits.
function matchesVisaAutomaton(lineText: string): boolean {
  let state = 0

  for (const char of lineText) {
    if (state === 0) {
      state = char === '4' ? 1 : -1
    } else if (!isAsciiDigit(char)) {
      state = -1
    } else if (state >= 1 && state < 13) {
      state += 1
    } else if (state === 13) {
      state = 14
    } else if (state === 14) {
      state = 15
    } else if (state === 15) {
      state = 16
    } else {
      state = -1
    }

    if (state === -1) {
      return false
    }
  }

  return state === 13 || state === 16
}

const REGEX_TASK_DEFINITIONS: Record<RegexSearchTaskId, RegexTaskDefinition> = {
  password: {
    id: 'password',
    titleKey: 'regex.task.password',
    matchesLine: lineText => PASSWORD_PATTERN.test(lineText),
  },
  visa: {
    id: 'visa',
    titleKey: 'regex.task.visa',
    matchesLine: lineText => VISA_PATTERN.test(lineText),
  },
  visaAutomaton: {
    id: 'visaAutomaton',
    titleKey: 'regex.task.visaAutomaton',
    matchesLine: matchesVisaAutomaton,
  },
  longitude: {
    id: 'longitude',
    titleKey: 'regex.task.longitude',
    matchesLine: lineText => LONGITUDE_PATTERN.test(lineText),
  },
}

export const REGEX_SEARCH_TASKS: RegexSearchTaskOption[] = Object.values(REGEX_TASK_DEFINITIONS)

export function searchRegexMatches(text: string, taskId: RegexSearchTaskId): RegexSearchMatch[] {
  const config = REGEX_TASK_DEFINITIONS[taskId]
  const matches: RegexSearchMatch[] = []
  const lines = text.split('\n')
  let nextId = 1
  let offset = 0

  for (let index = 0; index < lines.length; index += 1) {
    const lineText = lines[index]

    if (config.matchesLine(lineText)) {
      const from = offset
      const to = from + lineText.length

      matches.push({
        id: nextId,
        value: lineText,
        line: index + 1,
        col: 1,
        length: lineText.length,
        from,
        to,
      })

      nextId += 1
    }

    offset += lineText.length + 1
  }

  return matches
}