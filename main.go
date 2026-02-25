package main

import (
	"fyne.io/fyne/v2/app"

	"compiler/ui"
)

func main() {
	a := app.New()
	w := ui.NewCompilerWindow(a)
	w.ShowAndRun()
}
