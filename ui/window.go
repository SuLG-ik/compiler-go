package ui

import (
	"fmt"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type CompilerWindow struct {
	win         fyne.Window
	registry    *ActionRegistry
	currentFile string
	editor      *widget.Entry
	output      *widget.Entry
	statusLabel *widget.Label
}

func NewCompilerWindow(a fyne.App) fyne.Window {
	cw := &CompilerWindow{
		registry: newActionRegistry(),
	}
	cw.win = a.NewWindow(Strings.AppTitle)
	cw.win.Resize(fyne.NewSize(1100, 700))

	cw.registerActions()
	cw.buildUI()
	cw.updateTitle()
	return cw.win
}

func (cw *CompilerWindow) registerActions() {
	noop := func() {}

	cw.registry.Register(&Action{
		ID: ActionNew, Label: Strings.Actions.New, Tooltip: Strings.Tooltips.New,
		Icon: theme.DocumentCreateIcon(), Shortcut: ctrl(fyne.KeyN), Handler: noop,
	})
	cw.registry.Register(&Action{
		ID: ActionOpen, Label: Strings.Actions.Open, Tooltip: Strings.Tooltips.Open,
		Icon: theme.FolderOpenIcon(), Shortcut: ctrl(fyne.KeyO), Handler: noop,
	})
	cw.registry.Register(&Action{
		ID: ActionSave, Label: Strings.Actions.Save, Tooltip: Strings.Tooltips.Save,
		Icon: theme.DocumentSaveIcon(), Shortcut: ctrl(fyne.KeyS), Handler: noop,
	})
	cw.registry.Register(&Action{
		ID: ActionSaveAs, Label: Strings.Actions.SaveAs, Tooltip: Strings.Tooltips.SaveAs,
		Shortcut: ctrlShift(fyne.KeyS), Handler: noop,
	})
	cw.registry.Register(&Action{
		ID: ActionExit, Label: Strings.Actions.Exit, Tooltip: Strings.Tooltips.Exit,
		Handler: noop, IsQuit: true,
	})
	cw.registry.Register(&Action{
		ID: ActionUndo, Label: Strings.Actions.Undo, Tooltip: Strings.Tooltips.Undo,
		Icon: theme.ContentUndoIcon(), Shortcut: ctrl(fyne.KeyZ), Handler: noop,
	})
	cw.registry.Register(&Action{
		ID: ActionRedo, Label: Strings.Actions.Redo, Tooltip: Strings.Tooltips.Redo,
		Icon: theme.ContentRedoIcon(), Shortcut: ctrl(fyne.KeyY), Handler: noop,
	})
	cw.registry.Register(&Action{
		ID: ActionCut, Label: Strings.Actions.Cut, Tooltip: Strings.Tooltips.Cut,
		Icon: theme.ContentCutIcon(), Shortcut: ctrl(fyne.KeyX), Handler: noop,
	})
	cw.registry.Register(&Action{
		ID: ActionCopy, Label: Strings.Actions.Copy, Tooltip: Strings.Tooltips.Copy,
		Icon: theme.ContentCopyIcon(), Shortcut: ctrl(fyne.KeyC), Handler: noop,
	})
	cw.registry.Register(&Action{
		ID: ActionPaste, Label: Strings.Actions.Paste, Tooltip: Strings.Tooltips.Paste,
		Icon: theme.ContentPasteIcon(), Shortcut: ctrl(fyne.KeyV), Handler: noop,
	})
	cw.registry.Register(&Action{
		ID: ActionDelete, Label: Strings.Actions.Delete, Tooltip: Strings.Tooltips.Delete,
		Handler: noop,
	})
	cw.registry.Register(&Action{
		ID: ActionSelectAll, Label: Strings.Actions.SelectAll, Tooltip: Strings.Tooltips.SelectAll,
		Shortcut: ctrl(fyne.KeyA), Handler: noop,
	})
	cw.registry.Register(&Action{ID: ActionTask, Label: Strings.Actions.Task, Handler: noop})
	cw.registry.Register(&Action{ID: ActionGrammar, Label: Strings.Actions.Grammar, Handler: noop})
	cw.registry.Register(&Action{ID: ActionClass, Label: Strings.Actions.Class, Handler: noop})
	cw.registry.Register(&Action{ID: ActionMethod, Label: Strings.Actions.Method, Handler: noop})
	cw.registry.Register(&Action{ID: ActionTestEx, Label: Strings.Actions.TestEx, Handler: noop})
	cw.registry.Register(&Action{ID: ActionRefs, Label: Strings.Actions.References, Handler: noop})
	cw.registry.Register(&Action{ID: ActionSrcCode, Label: Strings.Actions.SourceCode, Handler: noop})
	cw.registry.Register(&Action{
		ID: ActionRun, Label: Strings.Actions.Run, Tooltip: Strings.Tooltips.Run,
		Icon: theme.MediaPlayIcon(), Shortcut: ctrl(fyne.KeyR), Handler: noop,
	})
	cw.registry.Register(&Action{
		ID: ActionHelp, Label: Strings.Actions.Help, Tooltip: Strings.Tooltips.Help,
		Icon: theme.HelpIcon(), Shortcut: hotkey(fyne.KeyF1), Handler: noop,
	})
	cw.registry.Register(&Action{
		ID: ActionAbout, Label: Strings.Actions.About, Tooltip: Strings.Tooltips.About,
		Icon: theme.InfoIcon(), Handler: noop,
	})
}

func (cw *CompilerWindow) buildUI() {
	var split *container.Split
	var statusBar fyne.CanvasObject
	cw.editor, cw.output, split = buildEditorPane()
	cw.statusLabel, statusBar = buildStatusBar()
	toolbar := buildToolbar(cw.registry)

	cw.win.SetContent(container.NewBorder(toolbar, statusBar, nil, nil, split))
	cw.win.SetMainMenu(buildMainMenu(cw.registry))
}

func (cw *CompilerWindow) updateTitle() {
	name := Strings.Untitled
	if cw.currentFile != "" {
		name = filepath.Base(cw.currentFile)
	}
	cw.win.SetTitle(fmt.Sprintf("%s \u2014 %s", name, Strings.AppTitle))
}
