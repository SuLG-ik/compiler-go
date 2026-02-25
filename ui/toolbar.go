package ui

import (
	"fyne.io/fyne/v2/widget"
)

func buildToolbar(r *ActionRegistry) *widget.Toolbar {
	act := func(id ActionID) *widget.ToolbarAction {
		a := r.Get(id)
		return widget.NewToolbarAction(a.Icon, a.Handler)
	}

	return widget.NewToolbar(
		act(ActionNew),
		act(ActionOpen),
		act(ActionSave),
		widget.NewToolbarSeparator(),
		act(ActionUndo),
		act(ActionRedo),
		act(ActionCopy),
		act(ActionCut),
		act(ActionPaste),
		act(ActionRun),
		act(ActionHelp),
		act(ActionAbout),
	)
}
