export type RegexSearchTaskId = 'password' | 'visa' | 'longitude'

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
  pattern: RegExp
}

const REGEX_TASK_DEFINITIONS: Record<RegexSearchTaskId, RegexTaskDefinition> = {
  password: {
    id: 'password',
    titleKey: 'regex.task.password',
    pattern: /^[А-Яа-яЁё0-9!@#$%^&*()_\-=+]{8,}$/u,
  },
  visa: {
    id: 'visa',
    titleKey: 'regex.task.visa',
    pattern: /^4[0-9]{12}([0-9]{3})?$/,
  },
  longitude: {
    id: 'longitude',
    titleKey: 'regex.task.longitude',
    pattern: /^(([0-9]{1,2}|1[0-7][0-9])°[0-5][0-9]'[0-5][0-9]"[EW]|180°00'00"[EW])$/u,
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

    if (config.pattern.test(lineText)) {
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