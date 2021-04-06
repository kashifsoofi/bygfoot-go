package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/kashifsoofi/bygfoot-go/store"
)

type App struct {
	store  store.Store
	app    fyne.App
	window fyne.Window
}

func NewApp(store store.Store) *App {
	app := App{
		store: store,
		app:   app.NewWithID("bygfoot"),
	}
	app.window = NewStartupWindow(app.app, store.Region())

	return &app
}

func (a *App) Run() {
	a.window.ShowAndRun()
}
