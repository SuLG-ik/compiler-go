package ui

var Strings = struct {
	AppTitle          string
	FileFilter        string
	Untitled          string
	OutputPlaceholder string
	EditorPlaceholder string

	Actions struct {
		New        string
		Open       string
		Save       string
		SaveAs     string
		Exit       string
		Undo       string
		Redo       string
		Cut        string
		Copy       string
		Paste      string
		Delete     string
		SelectAll  string
		Task       string
		Grammar    string
		Class      string
		Method     string
		TestEx     string
		References string
		SourceCode string
		Run        string
		Help       string
		About      string
	}

	Tooltips struct {
		New       string
		Open      string
		Save      string
		SaveAs    string
		Exit      string
		Undo      string
		Redo      string
		Cut       string
		Copy      string
		Paste     string
		Delete    string
		SelectAll string
		Run       string
		Help      string
		About     string
	}

	Menus struct {
		File string
		Edit string
		Text string
		Run  string
		Help string
	}

	Toolbar struct {
		Name string
	}

	Dialogs struct {
		UnsavedTitle   string
		UnsavedMessage string
		Save           string
		Discard        string
		Cancel         string
		Close          string
		AboutTitle     string
		AboutFmt       string // format: name \n version %s \n author
		HelpTitle      string
		HelpContent    string
	}

	Status struct {
		Ready       string
		NewDocument string
		OpenedFmt   string // format: "Открыт: %s"
		SavedFmt    string // format: "Сохранён: %s"
		Modified    string
		CursorFmt   string // format: "Стр. %d  Стб. %d"
	}
}{
	AppTitle:          "Compiler",
	FileFilter:        "*.txt",
	Untitled:          "Без названия",
	OutputPlaceholder: "Результаты работы языкового процессора",
	EditorPlaceholder: "Код здесь...",

	Actions: struct {
		New        string
		Open       string
		Save       string
		SaveAs     string
		Exit       string
		Undo       string
		Redo       string
		Cut        string
		Copy       string
		Paste      string
		Delete     string
		SelectAll  string
		Task       string
		Grammar    string
		Class      string
		Method     string
		TestEx     string
		References string
		SourceCode string
		Run        string
		Help       string
		About      string
	}{
		New:        "Создать",
		Open:       "Открыть",
		Save:       "Сохранить",
		SaveAs:     "Сохранить как",
		Exit:       "Выход",
		Undo:       "Отменить",
		Redo:       "Повторить",
		Cut:        "Вырезать",
		Copy:       "Копировать",
		Paste:      "Вставить",
		Delete:     "Удалить",
		SelectAll:  "Выделить все",
		Task:       "Постановка задачи",
		Grammar:    "Грамматика",
		Class:      "Классификация грамматики",
		Method:     "Метод анализа",
		TestEx:     "Тестовый пример",
		References: "Список литературы",
		SourceCode: "Исходный код программы",
		Run:        "Пуск",
		Help:       "Вызов справки",
		About:      "О программе",
	},

	Tooltips: struct {
		New       string
		Open      string
		Save      string
		SaveAs    string
		Exit      string
		Undo      string
		Redo      string
		Cut       string
		Copy      string
		Paste     string
		Delete    string
		SelectAll string
		Run       string
		Help      string
		About     string
	}{
		New:       "Создать новый документ",
		Open:      "Открыть документ",
		Save:      "Сохранить документ",
		SaveAs:    "Сохранить документ как",
		Exit:      "Выход из программы",
		Undo:      "Отменить последнее действие",
		Redo:      "Повторить последнее действие",
		Cut:       "Вырезать выделенный фрагмент",
		Copy:      "Копировать выделенный фрагмент",
		Paste:     "Вставить из буфера обмена",
		Delete:    "Удалить выделенный фрагмент",
		SelectAll: "Выделить весь текст",
		Run:       "Запустить синтаксический анализатор",
		Help:      "Открыть справку",
		About:     "Информация о программе",
	},

	Menus: struct {
		File string
		Edit string
		Text string
		Run  string
		Help string
	}{
		File: "Файл",
		Edit: "Правка",
		Text: "Текст",
		Run:  "Пуск",
		Help: "Справка",
	},

	Toolbar: struct {
		Name string
	}{
		Name: "Панель инструментов",
	},

	Dialogs: struct {
		UnsavedTitle   string
		UnsavedMessage string
		Save           string
		Discard        string
		Cancel         string
		Close          string
		AboutTitle     string
		AboutFmt       string
		HelpTitle      string
		HelpContent    string
	}{
		UnsavedTitle:   "Несохранённые изменения",
		UnsavedMessage: "Файл содержит несохранённые изменения.\nСохранить перед продолжением?",
		Save:           "Сохранить",
		Discard:        "Не сохранять",
		Cancel:         "Отмена",
		Close:          "Закрыть",
		AboutTitle:     "О программе",
		AboutFmt: "Compiler — учебный текстовый редактор\n\nВерсия: %s\n\n" +
			"Разработан в рамках курса\n«Системное программирование».\n\n" +
			"Язык реализации: Go\nGUI-фреймворк: Fyne v2",
		HelpTitle: "Справка",
		HelpContent: `# Руководство пользователя

## Меню «Файл»

| Команда | Горячие клавиши | Описание |
|---|---|---|
| Создать | Ctrl+N | Создаёт новый пустой документ. Если текущий документ изменён — программа предложит его сохранить. |
| Открыть | Ctrl+O | Открывает диалог выбора файла. Если текущий документ изменён — программа предложит его сохранить. |
| Сохранить | Ctrl+S | Сохраняет текущий документ. Если файл ещё не сохранялся — открывает диалог «Сохранить как». |
| Сохранить как | Ctrl+Shift+S | Открывает диалог выбора места и имени для сохранения документа. |
| Выход | — | Завершает работу программы. Если текущий документ изменён — программа предложит его сохранить. |

## Меню «Правка»

| Команда | Горячие клавиши | Описание |
|---|---|---|
| Отменить | Ctrl+Z | Отменяет последнее изменение в редакторе. |
| Повторить | Ctrl+Y | Повторяет отменённое изменение. |
| Вырезать | Ctrl+X | Вырезает выделенный текст в буфер обмена. |
| Копировать | Ctrl+C | Копирует выделенный текст в буфер обмена. |
| Вставить | Ctrl+V | Вставляет содержимое буфера обмена в позицию курсора. |
| Удалить | Del | Удаляет выделенный фрагмент текста. |
| Выделить все | Ctrl+A | Выделяет весь текст в области редактирования. |

## Меню «Пуск»

| Команда | Горячие клавиши | Описание |
|---|---|---|
| Пуск | Ctrl+R | Запускает синтаксический анализатор. Результаты отображаются в нижней области. |

## Меню «Справка»

| Команда | Горячие клавиши | Описание |
|---|---|---|
| Вызов справки | F1 | Открывает данное окно справки. |
| О программе | — | Отображает информацию о приложении и его версии. |

## Интерфейс

- **Область редактирования** — верхняя зона окна, предназначена для ввода и редактирования текста.
- **Область результатов** — нижняя зона окна (только чтение), предназначена для вывода результатов работы языкового процессора.
- **Разделитель** можно перетаскивать для изменения соотношения высот областей.
- **Строка состояния** — отображает положение курсора и текущий статус.
`,
	},

	Status: struct {
		Ready       string
		NewDocument string
		OpenedFmt   string
		SavedFmt    string
		Modified    string
		CursorFmt   string
	}{
		Ready:       "Готово",
		NewDocument: "Новый документ создан",
		OpenedFmt:   "Открыт: %s",
		SavedFmt:    "Сохранён: %s",
		Modified:    "Изменён",
		CursorFmt:   "Стр. %d  Стб. %d",
	},
}
