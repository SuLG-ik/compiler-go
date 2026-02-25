package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func buildStatusBar() (label *widget.Label, bar fyne.CanvasObject) {
	label = widget.NewLabel("")
	sep := canvas.NewRectangle(color.RGBA{R: 0xCC, G: 0xCC, B: 0xCC, A: 0xFF})
	sep.SetMinSize(fyne.NewSize(0, 1))
	bar = container.NewBorder(sep, nil, nil, nil, label)
	return
}
