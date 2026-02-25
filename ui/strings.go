package ui

var Strings = struct {
	AppTitle           string
	FileFilter         string
	Untitled           string
	OutputPlaceholder  string
	EditorPlaceholder  string

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
}
