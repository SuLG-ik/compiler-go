package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

type ActionID string

const (
	ActionNew       ActionID = "file.new"
	ActionOpen      ActionID = "file.open"
	ActionSave      ActionID = "file.save"
	ActionSaveAs    ActionID = "file.save_as"
	ActionExit      ActionID = "file.exit"
	ActionUndo      ActionID = "edit.undo"
	ActionRedo      ActionID = "edit.redo"
	ActionCut       ActionID = "edit.cut"
	ActionCopy      ActionID = "edit.copy"
	ActionPaste     ActionID = "edit.paste"
	ActionDelete    ActionID = "edit.delete"
	ActionSelectAll ActionID = "edit.select_all"
	ActionTask      ActionID = "text.task"
	ActionGrammar   ActionID = "text.grammar"
	ActionClass     ActionID = "text.class"
	ActionMethod    ActionID = "text.method"
	ActionTestEx    ActionID = "text.test_example"
	ActionRefs      ActionID = "text.references"
	ActionSrcCode   ActionID = "text.source_code"
	ActionRun       ActionID = "run.run"
	ActionHelp      ActionID = "help.help"
	ActionAbout     ActionID = "help.about"
)

type Action struct {
	ID       ActionID
	Label    string
	Tooltip  string
	Icon     fyne.Resource
	Shortcut fyne.Shortcut
	Handler  func()
	IsQuit   bool
}

type ActionRegistry struct {
	m map[ActionID]*Action
}

func newActionRegistry() *ActionRegistry {
	return &ActionRegistry{m: make(map[ActionID]*Action)}
}

func (r *ActionRegistry) Register(a *Action) {
	r.m[a.ID] = a
}

func (r *ActionRegistry) Get(id ActionID) *Action {
	if a, ok := r.m[id]; ok {
		return a
	}
	return &Action{Label: string(id), Handler: func() {}}
}

func ctrl(key fyne.KeyName) fyne.Shortcut {
	return &desktop.CustomShortcut{KeyName: key, Modifier: fyne.KeyModifierControl}
}

func ctrlShift(key fyne.KeyName) fyne.Shortcut {
	return &desktop.CustomShortcut{KeyName: key, Modifier: fyne.KeyModifierControl | fyne.KeyModifierShift}
}

func hotkey(key fyne.KeyName) fyne.Shortcut {
	return &desktop.CustomShortcut{KeyName: key}
}
