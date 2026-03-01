package main

import "context"

var Version = "dev"

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return new(App)
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetVersion() string {
	return Version
}
