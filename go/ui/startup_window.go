package ui

import (
	"fyne.io/fyne/theme"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewStartupWindow(a fyne.App) fyne.Window {
	w := a.NewWindow("Bygfoot Football Manager")

	widgetStartLeague := widget.NewSelect(getLeagues(""), func(name string) {
	})

	content := container.NewVBox(
		container.NewVBox(
			widget.NewLabel("Choose country"),
			widget.NewSelect(getCountries(), func(name string) {
			}),
		),
		widget.NewLabel("Choose team"),
		container.NewHSplit(
			container.NewScroll(
				// treeview_users
				widget.NewTreeWithStrings(getDemoTreeData()),
			),
			container.NewScroll(
				// treeview_startup
				widget.NewTreeWithStrings(getDemoTreeData()),
			),
		),
		widget.NewSeparator(),
		widget.NewLabel("Choose league to start in"),
		widgetStartLeague,
		widget.NewSeparator(),
		widget.NewLabel("Choose username"),
		widget.NewEntry(),
		widget.NewButtonWithIcon("Add user", theme.ContentAddIcon(), func() {}),
		container.NewHBox(
			widget.NewRadioGroup([]string{"Load team definitions", "Only names", "Don't load definitions"}, func(string) {}),
			widget.NewSeparator(),
			widget.NewCheck("Randomise teams in cups", func(bool) {}),
		),
		widget.NewSeparator(),
		container.NewHBox(
			widget.NewButtonWithIcon("Start", theme.ConfirmIcon(), func() {}),
			widget.NewButtonWithIcon("Back to splash", theme.NavigateBackIcon(), func() {}),
			widget.NewButton("Quit", func() {}),
		),
	)

	w.SetContent(content)
	w.CenterOnScreen()
	return w
}

func getCountries() []string {
	return []string{
		"Pakistan",
		"England",
		"France",
		"Germany",
	}
}

func getLeagues(country string) []string {
	return []string{
		country + " L1",
		country + " L2",
		country + " L3",
	}
}

func getDemoTreeData() map[string][]string {
	return map[string][]string{
		"":  {"A"},
		"A": {"B", "D", "H", "J", "L", "O", "P", "S", "V"},
		"B": {"C"},
		"C": {"abc"},
		"D": {"E"},
		"E": {"F", "G"},
		"F": {"adef"},
		"G": {"adeg"},
		"H": {"I"},
		"I": {"ahi"},
		"O": {"ao"},
		"P": {"Q"},
		"Q": {"R"},
		"R": {"apqr"},
		"S": {"T"},
		"T": {"U"},
		"U": {"astu"},
		"V": {"W"},
		"W": {"X"},
		"X": {"Y"},
		"Y": {"Z"},
		"Z": {"avwxyz"},
	}
}
