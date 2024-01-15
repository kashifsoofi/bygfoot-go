package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func NewSplashWindow(a fyne.App) fyne.Window {
	w := a.NewWindow("Bygfoot Football Manager")

	content := container.NewVBox(
		container.NewAppTabs(
			container.NewTabItem(
				"    ",
				widget.NewLabel("TODO: Load image"),
			),
			container.NewTabItem(
				"Contributors",
				container.NewScroll(widget.NewLabel("Contributors List")),
			),
		),
		widget.NewSeparator(),
		// TODO: Horizontal Separator
		container.NewVBox(
			container.NewHBox(
				widget.NewIcon(theme.InfoIcon()),
				widget.NewLabel("Did you know?"),
				widget.NewLabel("0 / 15"),
			),
			widget.NewLabel("Hint 1"),
			container.NewHBox(
				widget.NewButtonWithIcon("Previous", theme.NavigateBackIcon(), func() {}),
				widget.NewButtonWithIcon("Next", theme.NavigateNextIcon(), func() {}),
			),
		),
		// TODO: Horizontal Separator
		container.NewHBox(
			widget.NewButtonWithIcon("Start new game", theme.FolderNewIcon(), func() {}),
			widget.NewButtonWithIcon("Load game", theme.FolderOpenIcon(), func() {}),
			widget.NewButtonWithIcon("Resume game", theme.ViewRestoreIcon(), func() {}),
			widget.NewButtonWithIcon("Quit", theme.CancelIcon(), func() {}),
		),
		container.NewHBox(
			widget.NewLabel("Ready"),
			widget.NewProgressBar(),
		),
	)

	w.SetContent(content)
	w.CenterOnScreen()
	return w
}
