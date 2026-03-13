package main

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var Version = "dev"

type App struct {
	ctx     context.Context
	canQuit bool
}

func NewApp() *App {
	return new(App)
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// beforeClose вызывается при нажатии на крестик окна.
// Возвращает true — предотвращает закрытие (фронтенд сам решит).
func (a *App) beforeClose(ctx context.Context) bool {
	if a.canQuit {
		return false // разрешить закрытие
	}
	runtime.EventsEmit(ctx, "app:beforeclose")
	return true // заблокировать закрытие
}

// AllowQuit вызывается из фронтенда, когда всё сохранено.
func (a *App) AllowQuit() {
	a.canQuit = true
	runtime.Quit(a.ctx)
}

func (a *App) GetVersion() string {
	return Version
}
