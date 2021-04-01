package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/kashifsoofi/bygfoot-go/ui"
)

func main() {
	a := app.New()
	w := ui.NewSplashWindow(a)

	w.ShowAndRun()
}
