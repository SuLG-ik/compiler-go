package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func (cw *CompilerWindow) doHelp() {
	rt := widget.NewRichTextFromMarkdown(Strings.Dialogs.HelpContent)
	rt.Wrapping = fyne.TextWrapWord

	scroll := container.NewScroll(rt)

	w := fyne.CurrentApp().NewWindow(Strings.Dialogs.HelpTitle)
	w.Resize(fyne.NewSize(700, 500))
	w.SetContent(scroll)
	w.CenterOnScreen()
	w.Show()
}

func (cw *CompilerWindow) doAbout() {
	msg := fmt.Sprintf(Strings.Dialogs.AboutFmt, cw.version)
	dialog.ShowInformation(Strings.Dialogs.AboutTitle, msg, cw.win)
}

func showInfoDialog(title, body string, parent fyne.Window) {
	lbl := widget.NewLabel(body)
	lbl.Wrapping = fyne.TextWrapWord

	scroll := container.NewScroll(lbl)
	scroll.SetMinSize(fyne.NewSize(480, 300))

	d := dialog.NewCustom(title, Strings.Dialogs.Close, scroll, parent)
	d.Resize(fyne.NewSize(540, 380))
	d.Show()
}
