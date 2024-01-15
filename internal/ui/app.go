package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/kashifsoofi/bygfoot-go/internal/store"

	log "github.com/sirupsen/logrus"
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
	app.window = NewSplashWindow(app.app)
	//app.window = NewStartupWindow(app.app, store.Region(), store.League())

	return &app
}

func (a *App) Run() {
	log.Debug("app.Run")
	a.window.ShowAndRun()
}
