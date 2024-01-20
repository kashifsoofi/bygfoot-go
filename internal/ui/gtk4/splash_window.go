package gtk4

import (
	_ "embed"
	"fmt"

	"log/slog"

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
	hints      []string
}

func NewSplashWindow(app *gtk.Application, hintsStore store.HintsStore) *SplashWindow {
	// You can build UIs using Cambalache (https://flathub.org/apps/details/ar.xjuan.Cambalache)
	builder := gtk.NewBuilderFromString(splashXml, len(splashXml))

	sw := &SplashWindow{hintsStore: hintsStore}

	// MainWindow and Button are object IDs from the UI file
	sw.window = builder.GetObject("SplashWindow").Cast().(*gtk.Window)
	sw.window.ConnectCloseRequest(sw.splashWindowCloseRequestHandler)

	sw.labelSplashHintCounter = builder.GetObject("label_splash_hintcounter").Cast().(*gtk.Label)
	sw.labelSplashHint = builder.GetObject("label_splash_hint").Cast().(*gtk.Label)
	sw.buttonSplashHintBack = builder.GetObject("button_splash_hint_back").Cast().(*gtk.Button)
	sw.buttonSplashHintBack.ConnectClicked(sw.buttonSplashHintBackClicked)

	sw.buttonSplashHintNext = builder.GetObject("button_splash_hint_next").Cast().(*gtk.Button)
	sw.buttonSplashHintNext.ConnectClicked(sw.buttonSplashHintNextClicked)

	buttonSplashNewGame := builder.GetObject("button_splash_new_game").Cast().(*gtk.Button)
	buttonSplashNewGame.ConnectClicked(sw.buttonSplashNewGameClicked)

	buttonSplashLoadGame := builder.GetObject("button_splash_load_game").Cast().(*gtk.Button)
	buttonSplashLoadGame.ConnectClicked(sw.buttonSplashLoadGameClicked)

	buttonSplashResumeGame := builder.GetObject("button_splash_resume_game").Cast().(*gtk.Button)
	buttonSplashResumeGame.ConnectClicked(sw.buttonSplashResumeGameClicked)

	buttonSplashQuit := builder.GetObject("button_splash_quit").Cast().(*gtk.Button)
	buttonSplashQuit.ConnectClicked(sw.buttonSplashQuitClicked)

	sw.hintNum = hintsStore.LoadHintNumber()
	sw.hints = hintsStore.Hints()
	sw.showHint()

	app.AddWindow(sw.window)
	return sw
}

func (w *SplashWindow) Show() {
	w.window.Show()
}

func (w *SplashWindow) splashWindowDestroy() {
	slog.Debug("splashWindowDestroy")
	w.hintsStore.SaveHintNumber(w.hintNum)
	w.window.Destroy()
}

func (w *SplashWindow) splashWindowCloseRequestHandler() bool {
	slog.Debug("splashWindowCloseRequestHandler")
	w.splashWindowDestroy()
	return false
}

func (w *SplashWindow) showHint() {
	totalHints := len(w.hints) - 1
	if w.hintNum < 0 {
		w.hintNum = totalHints
	}
	if w.hintNum > totalHints {
		w.hintNum = 0
	}

	hint := w.hints[w.hintNum]
	w.labelSplashHint.SetLabel(hint)

	hintCounter := fmt.Sprintf("(%d/%d)", w.hintNum+1, totalHints+1)
	w.labelSplashHintCounter.SetLabel(hintCounter)
}

func (w *SplashWindow) buttonSplashHintBackClicked() {
	slog.Debug("buttonSplashHintBackClicked")
	w.hintNum -= 1
	w.showHint()
}

func (w *SplashWindow) buttonSplashHintNextClicked() {
	slog.Debug("buttonSplashHintNextClicked")
	w.hintNum += 1
	w.showHint()
}

func (w *SplashWindow) buttonSplashNewGameClicked() {
	slog.Debug("buttonSplashNewGameClicked")
}

func (w *SplashWindow) buttonSplashLoadGameClicked() {
	slog.Debug("buttonSplashLoadGameClicked")
}

func (w *SplashWindow) buttonSplashResumeGameClicked() {
	slog.Debug("buttonSplashResumeGameClicked")
}

func (w *SplashWindow) buttonSplashQuitClicked() {
	slog.Debug("buttonSplashQuitClicked")
	app := w.window.Application()
	w.splashWindowDestroy()
	app.Quit()
}
