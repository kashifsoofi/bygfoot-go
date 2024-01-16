package gtk4

import (
	_ "embed"
	"strconv"

	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

//go:embed splash.ui
var splashXml string

func NewSplashWindow(app *gtk.Application) *gtk.Window {
	// You can build UIs using Cambalache (https://flathub.org/apps/details/ar.xjuan.Cambalache)
	builder := gtk.NewBuilderFromString(splashXml, len(splashXml))

	// MainWindow and Button are object IDs from the UI file
	window := builder.GetObject("SplashWindow").Cast().(*gtk.Window)
	button := builder.GetObject("Button").Cast().(*gtk.Button)
	counter := 0
	button.Connect("clicked", func() {
		button.SetLabel(strconv.Itoa(counter))
		counter++
	})

	app.AddWindow(window)
	return window
}
