package gtk4

import (
	_ "embed"
	"os"

	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/kashifsoofi/bygfoot-go/internal/store"
	log "github.com/sirupsen/logrus"
)

type App struct {
	store store.Store
	app   *gtk.Application
	// window fyne.Window
}

const appId = "com.github.kashifsoofi.bygfoot-go"

func NewApp(store store.Store) *App {
	gtkApp := gtk.NewApplication(appId, gio.ApplicationFlagsNone)
	gtkApp.ConnectActivate(func() { activate(gtkApp) })

	app := App{
		store: store,
		app:   gtkApp,
	}

	return &app
}

func (a *App) Run() {
	log.Debug("app.Run")
	if code := a.app.Run(os.Args); code > 0 {
		os.Exit(code)
	}
}

func activate(app *gtk.Application) {
	window := NewSplashWindow(app)
	window.Show()
}
