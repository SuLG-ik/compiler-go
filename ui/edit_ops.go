package ui

import "fyne.io/fyne/v2"

// focusEditor переводит фокус на область редактирования.
func (cw *CompilerWindow) focusEditor() {
	cw.win.Canvas().Focus(cw.editor)
}

// doUndo отменяет последнее изменение в редакторе (встроенный стек Fyne ≥ 2.5).
func (cw *CompilerWindow) doUndo() {
	cw.focusEditor()
	cw.editor.Undo()
}

// doRedo повторяет отменённое изменение (встроенный стек Fyne ≥ 2.5).
func (cw *CompilerWindow) doRedo() {
	cw.focusEditor()
	cw.editor.Redo()
}

// doCut вырезает выделенный фрагмент в системный буфер обмена.
func (cw *CompilerWindow) doCut() {
	cw.focusEditor()
	cw.editor.TypedShortcut(&fyne.ShortcutCut{Clipboard: cw.win.Clipboard()})
}

// doCopy копирует выделенный фрагмент в системный буфер обмена.
func (cw *CompilerWindow) doCopy() {
	cw.focusEditor()
	cw.editor.TypedShortcut(&fyne.ShortcutCopy{Clipboard: cw.win.Clipboard()})
}

// doPaste вставляет содержимое системного буфера обмена в позицию курсора.
func (cw *CompilerWindow) doPaste() {
	cw.focusEditor()
	cw.editor.TypedShortcut(&fyne.ShortcutPaste{Clipboard: cw.win.Clipboard()})
}

// doDelete удаляет выделенный фрагмент (или символ после курсора, если выделения нет).
func (cw *CompilerWindow) doDelete() {
	cw.focusEditor()
	cw.editor.TypedKey(&fyne.KeyEvent{Name: fyne.KeyDelete})
}

// doSelectAll выделяет весь текст в редакторе.
func (cw *CompilerWindow) doSelectAll() {
	cw.focusEditor()
	cw.editor.TypedShortcut(&fyne.ShortcutSelectAll{})
}
