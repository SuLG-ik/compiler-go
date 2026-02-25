package main

import (
	"fyne.io/fyne/v2/app"

	"compiler/ui"
)

func main() {
	a := app.New()
	a.SetIcon(resourceIconPng)
	w := ui.NewCompilerWindow(a, Version)
	w.ShowAndRun()
}
