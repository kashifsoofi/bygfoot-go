package gtk4

import (
	_ "embed"
	"log/slog"
	"os"

	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/kashifsoofi/bygfoot-go/internal/store"
)

type App struct {
	store store.Store
	app   *gtk.Application
	// window fyne.Window
}

const appId = "com.github.kashifsoofi.bygfoot-go"

func NewApp(store store.Store) *App {
	app := App{store: store}
	app.app = gtk.NewApplication(appId, gio.ApplicationFlagsNone)
	app.app.ConnectActivate(app.activate)

	return &app
}

func (a *App) Run() {
	slog.Debug("app.Run")
	if code := a.app.Run(os.Args); code > 0 {
		os.Exit(code)
	}
}

func (a *App) activate() {
	window := NewSplashWindow(a.app, a.store.Hints())
	window.Show()
}
