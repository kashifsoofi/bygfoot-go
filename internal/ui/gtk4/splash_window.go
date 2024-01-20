package gtk4

import (
	_ "embed"
	"fmt"

	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/kashifsoofi/bygfoot-go/internal/store"
)

//go:embed splash.ui
var splashXml string

type SplashWindow struct {
	window                 *gtk.Window
	labelSplashHintCounter *gtk.Label
	labelSplashHint        *gtk.Label
	buttonSplashHintBack   *gtk.Button
	buttonSplashHintNext   *gtk.Button

	hintsStore store.HintsStore
	hintNum    int
}

func NewSplashWindow(app *gtk.Application, hintsStore store.HintsStore) *SplashWindow {
	// You can build UIs using Cambalache (https://flathub.org/apps/details/ar.xjuan.Cambalache)
	builder := gtk.NewBuilderFromString(splashXml, len(splashXml))

	sw := &SplashWindow{hintsStore: hintsStore}

	// MainWindow and Button are object IDs from the UI file
	sw.window = builder.GetObject("SplashWindow").Cast().(*gtk.Window)
	sw.labelSplashHintCounter = builder.GetObject("label_splash_hintcounter").Cast().(*gtk.Label)
	sw.labelSplashHint = builder.GetObject("label_splash_hint").Cast().(*gtk.Label)
	sw.buttonSplashHintBack = builder.GetObject("button_splash_hint_back").Cast().(*gtk.Button)
	sw.buttonSplashHintNext = builder.GetObject("button_splash_hint_next").Cast().(*gtk.Button)

	sw.buttonSplashHintBack.ConnectClicked(sw.buttonSplashHintBackClicked)
	sw.buttonSplashHintNext.ConnectClicked(sw.buttonSplashHintNextClicked)

	sw.hintNum = hintsStore.LoadHintNumber()
	sw.showHint()

	app.AddWindow(sw.window)
	return sw
}

func (w *SplashWindow) Show() {
	w.window.Show()
}

func (w *SplashWindow) showHint() {
	totalHints := w.hintsStore.TotalHints()
	if w.hintNum < 0 {
		w.hintNum = totalHints
	}
	if w.hintNum > totalHints {
		w.hintNum = 0
	}

	hint := w.hintsStore.Hint(w.hintNum)
	w.labelSplashHint.SetLabel(hint)

	hintCounter := fmt.Sprintf("(%d/%d)", w.hintNum+1, totalHints+1)
	w.labelSplashHintCounter.SetLabel(hintCounter)
}

func (w *SplashWindow) buttonSplashHintBackClicked() {
	w.hintNum -= 1
	w.showHint()
}

func (w *SplashWindow) buttonSplashHintNextClicked() {
	w.hintNum += 1
	w.showHint()
}
