package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/kashifsoofi/bygfoot-go/store"
	log "github.com/sirupsen/logrus"
)

func NewStartupWindow(a fyne.App, store store.RegionStore) fyne.Window {
	w := a.NewWindow("Bygfoot Football Manager")

	widgetChooseCountry := widget.NewSelect(getCountries(store), func(name string) {

	})

	widgetStartLeague := widget.NewSelect(getLeagues(""), func(name string) {
	})

	content := container.NewVBox(
		container.NewVBox(
			widget.NewLabel("Choose country"),
			widgetChooseCountry,
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

func getCountries(store store.RegionStore) []string {
	countries, err := store.ListCountries()
	if err != nil {
		log.WithFields(log.Fields{
			"Error": err,
		}).Error("failed to load countries")
	}

	countryNames := []string{}
	for _, c := range countries {
		countryNames = append(countryNames, c.Name)
	}
	return countryNames
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
