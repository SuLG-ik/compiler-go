import { createContext, useContext, useState, ReactNode, useCallback } from 'react'
import { ru } from './ru'
import { en } from './en'

export type Lang = 'ru' | 'en'

const dictionaries: Record<Lang, Record<string, string>> = { ru, en }

interface I18nContextValue {
  lang: Lang
  setLang: (lang: Lang) => void
  t: (key: string, params?: Record<string, string>) => string
}

const I18nContext = createContext<I18nContextValue>({
  lang: 'ru',
  setLang: () => {},
  t: (key: string) => key,
})

export function I18nProvider({ children }: { children: ReactNode }) {
  const [lang, setLang] = useState<Lang>(() => {
    const saved = localStorage.getItem('app-lang')
    return (saved === 'en' ? 'en' : 'ru') as Lang
  })

  const changeLang = useCallback((newLang: Lang) => {
    setLang(newLang)
    localStorage.setItem('app-lang', newLang)
  }, [])

  const t = useCallback((key: string, params?: Record<string, string>): string => {
    let str = dictionaries[lang][key] ?? key
    if (params) {
      Object.entries(params).forEach(([k, v]) => {
        str = str.replace(new RegExp(`\\{${k}\\}`, 'g'), v)
      })
    }
    return str
  }, [lang])

  return (
    <I18nContext.Provider value={{ lang, setLang: changeLang, t }}>
      {children}
    </I18nContext.Provider>
  )
}

export function useTranslation() {
  return useContext(I18nContext)
}
