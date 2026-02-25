package ui

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

// confirmUnsaved спрашивает пользователя о сохранении, если текст изменён,
// после чего вызывает then(). Если изменений нет — then() вызывается сразу.
func (cw *CompilerWindow) confirmUnsaved(then func()) {
	if !cw.dirty {
		then()
		return
	}

	var d dialog.Dialog

	saveBtn := widget.NewButton(Strings.Dialogs.Save, func() {
		d.Hide()
		cw.doSave(then)
	})
	saveBtn.Importance = widget.HighImportance

	discardBtn := widget.NewButton(Strings.Dialogs.Discard, func() {
		d.Hide()
		then()
	})

	content := container.NewHBox(saveBtn, discardBtn)
	d = dialog.NewCustom(
		Strings.Dialogs.UnsavedTitle,
		Strings.Dialogs.Cancel,
		content,
		cw.win,
	)
	d.Show()
}

// doNew очищает редактор и сбрасывает текущий файл.
func (cw *CompilerWindow) doNew() {
	cw.confirmUnsaved(func() {
		cw.editor.SetText("")
		cw.currentFile = ""
		cw.dirty = false
		cw.updateTitle()
		cw.setStatus(Strings.Status.NewDocument)
	})
}

// doOpen открывает диалог выбора файла и загружает его содержимое в редактор.
func (cw *CompilerWindow) doOpen() {
	cw.confirmUnsaved(func() {
		fd := dialog.NewFileOpen(func(r fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, cw.win)
				return
			}
			if r == nil {
				return
			}
			defer r.Close()

			data, err := io.ReadAll(r)
			if err != nil {
				dialog.ShowError(err, cw.win)
				return
			}

			cw.editor.SetText(string(data))
			cw.currentFile = r.URI().Path()
			cw.dirty = false
			cw.updateTitle()
			cw.setStatus(fmt.Sprintf(Strings.Status.OpenedFmt, filepath.Base(cw.currentFile)))
		}, cw.win)

		fd.SetFilter(storage.NewExtensionFileFilter([]string{
			".txt", ".go", ".c", ".cpp", ".h", ".pas", ".cs", ".py",
		}))
		fd.Show()
	})
}

// doSave сохраняет документ по текущему пути. Если путь не задан — открывает «Сохранить как».
func (cw *CompilerWindow) doSave(done func()) {
	if cw.currentFile == "" {
		cw.doSaveAs(done)
		return
	}
	cw.writeFile(cw.currentFile, done)
}

// doSaveAs открывает диалог выбора места сохранения.
func (cw *CompilerWindow) doSaveAs(done func()) {
	fd := dialog.NewFileSave(func(w fyne.URIWriteCloser, err error) {
		if err != nil {
			dialog.ShowError(err, cw.win)
			return
		}
		if w == nil {
			return
		}
		defer w.Close()

		path := w.URI().Path()
		if err := os.WriteFile(path, []byte(cw.editor.Text), 0644); err != nil {
			dialog.ShowError(err, cw.win)
			return
		}
		cw.currentFile = path
		cw.dirty = false
		cw.updateTitle()
		cw.setStatus(fmt.Sprintf(Strings.Status.SavedFmt, filepath.Base(path)))

		if done != nil {
			done()
		}
	}, cw.win)

	name := "untitled.txt"
	if cw.currentFile != "" {
		name = filepath.Base(cw.currentFile)
	}
	fd.SetFileName(name)
	fd.SetFilter(storage.NewExtensionFileFilter([]string{".txt", ".go", ".c", ".cpp", ".h", ".pas"}))
	fd.Show()
}

// writeFile записывает содержимое редактора в файл по указанному пути.
func (cw *CompilerWindow) writeFile(path string, done func()) {
	if err := os.WriteFile(path, []byte(cw.editor.Text), 0644); err != nil {
		dialog.ShowError(err, cw.win)
		return
	}
	cw.dirty = false
	cw.updateTitle()
	cw.setStatus(fmt.Sprintf(Strings.Status.SavedFmt, filepath.Base(path)))

	if done != nil {
		done()
	}
}

// doExit спрашивает о сохранении изменений и завершает работу.
func (cw *CompilerWindow) doExit() {
	cw.confirmUnsaved(func() {
		cw.win.Close()
	})
}
