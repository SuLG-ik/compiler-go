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
	version     string
	currentFile string
	dirty       bool
	editor      *widget.Entry
	output      *widget.Entry
	statusLabel *widget.Label
}

func NewCompilerWindow(a fyne.App, version string) fyne.Window {
	cw := &CompilerWindow{
		registry: newActionRegistry(),
		version:  version,
	}
	cw.win = a.NewWindow(Strings.AppTitle)
	cw.win.Resize(fyne.NewSize(1100, 700))

	cw.registerActions()
	cw.buildUI()
	cw.updateTitle()
	return cw.win
}

func (cw *CompilerWindow) registerActions() {
	cw.registry.Register(&Action{
		ID: ActionNew, Label: Strings.Actions.New, Tooltip: Strings.Tooltips.New,
		Icon: theme.DocumentCreateIcon(), Shortcut: ctrl(fyne.KeyN), Handler: cw.doNew,
	})
	cw.registry.Register(&Action{
		ID: ActionOpen, Label: Strings.Actions.Open, Tooltip: Strings.Tooltips.Open,
		Icon: theme.FolderOpenIcon(), Shortcut: ctrl(fyne.KeyO), Handler: cw.doOpen,
	})
	cw.registry.Register(&Action{
		ID: ActionSave, Label: Strings.Actions.Save, Tooltip: Strings.Tooltips.Save,
		Icon: theme.DocumentSaveIcon(), Shortcut: ctrl(fyne.KeyS),
		Handler: func() { cw.doSave(nil) },
	})
	cw.registry.Register(&Action{
		ID: ActionSaveAs, Label: Strings.Actions.SaveAs, Tooltip: Strings.Tooltips.SaveAs,
		Shortcut: ctrlShift(fyne.KeyS), Handler: func() { cw.doSaveAs(nil) },
	})
	cw.registry.Register(&Action{
		ID: ActionExit, Label: Strings.Actions.Exit, Tooltip: Strings.Tooltips.Exit,
		Handler: cw.doExit, IsQuit: true,
	})
	cw.registry.Register(&Action{
		ID: ActionUndo, Label: Strings.Actions.Undo, Tooltip: Strings.Tooltips.Undo,
		Icon: theme.ContentUndoIcon(), Shortcut: ctrl(fyne.KeyZ), Handler: cw.doUndo,
	})
	cw.registry.Register(&Action{
		ID: ActionRedo, Label: Strings.Actions.Redo, Tooltip: Strings.Tooltips.Redo,
		Icon: theme.ContentRedoIcon(), Shortcut: ctrl(fyne.KeyY), Handler: cw.doRedo,
	})
	cw.registry.Register(&Action{
		ID: ActionCut, Label: Strings.Actions.Cut, Tooltip: Strings.Tooltips.Cut,
		Icon: theme.ContentCutIcon(), Shortcut: ctrl(fyne.KeyX), Handler: cw.doCut,
	})
	cw.registry.Register(&Action{
		ID: ActionCopy, Label: Strings.Actions.Copy, Tooltip: Strings.Tooltips.Copy,
		Icon: theme.ContentCopyIcon(), Shortcut: ctrl(fyne.KeyC), Handler: cw.doCopy,
	})
	cw.registry.Register(&Action{
		ID: ActionPaste, Label: Strings.Actions.Paste, Tooltip: Strings.Tooltips.Paste,
		Icon: theme.ContentPasteIcon(), Shortcut: ctrl(fyne.KeyV), Handler: cw.doPaste,
	})
	cw.registry.Register(&Action{
		ID: ActionDelete, Label: Strings.Actions.Delete, Tooltip: Strings.Tooltips.Delete,
		Handler: cw.doDelete,
	})
	cw.registry.Register(&Action{
		ID: ActionSelectAll, Label: Strings.Actions.SelectAll, Tooltip: Strings.Tooltips.SelectAll,
		Shortcut: ctrl(fyne.KeyA), Handler: cw.doSelectAll,
	})
	cw.registry.Register(&Action{ID: ActionTask, Label: Strings.Actions.Task, Handler: cw.doTask})
	cw.registry.Register(&Action{ID: ActionGrammar, Label: Strings.Actions.Grammar, Handler: cw.doGrammar})
	cw.registry.Register(&Action{ID: ActionClass, Label: Strings.Actions.Class, Handler: cw.doGrammarClass})
	cw.registry.Register(&Action{ID: ActionMethod, Label: Strings.Actions.Method, Handler: cw.doMethod})
	cw.registry.Register(&Action{ID: ActionTestEx, Label: Strings.Actions.TestEx, Handler: cw.doTestEx})
	cw.registry.Register(&Action{ID: ActionRefs, Label: Strings.Actions.References, Handler: cw.doRefs})
	cw.registry.Register(&Action{ID: ActionSrcCode, Label: Strings.Actions.SourceCode, Handler: cw.doSrcCode})
	cw.registry.Register(&Action{
		ID: ActionRun, Label: Strings.Actions.Run, Tooltip: Strings.Tooltips.Run,
		Icon: theme.MediaPlayIcon(), Shortcut: ctrl(fyne.KeyR),
		Handler: func() { cw.setStatus("Анализатор ещё не реализован") },
	})
	cw.registry.Register(&Action{
		ID: ActionHelp, Label: Strings.Actions.Help, Tooltip: Strings.Tooltips.Help,
		Icon: theme.HelpIcon(), Shortcut: hotkey(fyne.KeyF1), Handler: cw.doHelp,
	})
	cw.registry.Register(&Action{
		ID: ActionAbout, Label: Strings.Actions.About, Tooltip: Strings.Tooltips.About,
		Icon: theme.InfoIcon(), Handler: cw.doAbout,
	})
}

func (cw *CompilerWindow) buildUI() {
	var split *container.Split
	var statusBar fyne.CanvasObject
	cw.editor, cw.output, split = buildEditorPane()
	cw.statusLabel, statusBar = buildStatusBar()
	toolbar := buildToolbar(cw.registry)

	// Отслеживаем изменения текста → флаг dirty + строка состояния
	cw.editor.OnChanged = func(_ string) {
		cw.markDirty()
	}
	// Обновляем позицию курсора в строке состояния
	cw.editor.OnCursorChanged = cw.updateCursorStatus

	cw.win.SetContent(container.NewBorder(toolbar, statusBar, nil, nil, split))
	cw.win.SetMainMenu(buildMainMenu(cw.registry))

	// Перехватываем кнопку закрытия окна
	cw.win.SetCloseIntercept(cw.doExit)

	cw.setStatus(Strings.Status.Ready)
}

func (cw *CompilerWindow) updateTitle() {
	name := Strings.Untitled
	if cw.currentFile != "" {
		name = filepath.Base(cw.currentFile)
	}
	dirtyMark := ""
	if cw.dirty {
		dirtyMark = "* "
	}
	cw.win.SetTitle(fmt.Sprintf("%s%s \u2014 %s", dirtyMark, name, Strings.AppTitle))
}

// markDirty выставляет флаг несохранённых изменений и обновляет заголовок.
func (cw *CompilerWindow) markDirty() {
	if !cw.dirty {
		cw.dirty = true
		cw.updateTitle()
	}
	cw.setStatus(Strings.Status.Modified)
}

// setStatus обновляет текст строки состояния.
func (cw *CompilerWindow) setStatus(msg string) {
	cw.statusLabel.SetText(msg)
}

// updateCursorStatus обновляет строку состояния с текущей позицией курсора.
func (cw *CompilerWindow) updateCursorStatus() {
	row := cw.editor.CursorRow + 1
	col := cw.editor.CursorColumn + 1
	cw.statusLabel.SetText(fmt.Sprintf(Strings.Status.CursorFmt, row, col))
}
