package ui

import (
	"fyne.io/fyne/v2"
)

func buildMainMenu(r *ActionRegistry) *fyne.MainMenu {
	item := func(id ActionID) *fyne.MenuItem {
		a := r.Get(id)
		mi := fyne.NewMenuItem(a.Label, a.Handler)
		mi.Icon = a.Icon
		mi.Shortcut = a.Shortcut
		mi.IsQuit = a.IsQuit
		return mi
	}

	fileMenu := fyne.NewMenu(Strings.Menus.File,
		item(ActionNew),
		item(ActionOpen),
		item(ActionSave),
		item(ActionSaveAs),
		fyne.NewMenuItemSeparator(),
		item(ActionExit),
	)

	editMenu := fyne.NewMenu(Strings.Menus.Edit,
		item(ActionUndo),
		item(ActionRedo),
		fyne.NewMenuItemSeparator(),
		item(ActionCut),
		item(ActionCopy),
		item(ActionPaste),
		item(ActionDelete),
		fyne.NewMenuItemSeparator(),
		item(ActionSelectAll),
	)

	textMenu := fyne.NewMenu(Strings.Menus.Text,
		item(ActionTask),
		item(ActionGrammar),
		item(ActionClass),
		item(ActionMethod),
		item(ActionTestEx),
		item(ActionRefs),
		item(ActionSrcCode),
	)

	runMenu := fyne.NewMenu(Strings.Menus.Run,
		item(ActionRun),
	)

	helpMenu := fyne.NewMenu(Strings.Menus.Help,
		item(ActionHelp),
		item(ActionAbout),
	)

	return fyne.NewMainMenu(fileMenu, editMenu, textMenu, runMenu, helpMenu)
}
