package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func buildEditorPane() (editor *widget.Entry, output *widget.Entry, split *container.Split) {
	editor = widget.NewMultiLineEntry()
	editor.TextStyle = fyne.TextStyle{Monospace: true}
	editor.SetPlaceHolder(Strings.EditorPlaceholder)
	editor.Wrapping = fyne.TextWrapOff

	output = widget.NewMultiLineEntry()
	output.TextStyle = fyne.TextStyle{Monospace: true}
	output.SetPlaceHolder(Strings.OutputPlaceholder)
	output.Wrapping = fyne.TextWrapOff
	output.Disable()

	outputBg := canvas.NewRectangle(color.RGBA{R: 0xF5, G: 0xF5, B: 0xF5, A: 0xFF})
	outputPane := container.NewStack(outputBg, output)

	split = container.NewVSplit(editor, outputPane)
	split.SetOffset(0.75)
	return
}
