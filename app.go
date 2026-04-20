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

func (a *App) beforeClose(ctx context.Context) bool {
	if a.canQuit {
		return false
	}
	runtime.EventsEmit(ctx, "app:beforeclose")
	return true
}

func (a *App) AllowQuit() {
	a.canQuit = true
	runtime.Quit(a.ctx)
}

func (a *App) GetVersion() string {
	return Version
}
