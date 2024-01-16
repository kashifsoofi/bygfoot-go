package fyne

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/kashifsoofi/bygfoot-go/internal/store"
	log "github.com/sirupsen/logrus"
)

var countryNameToId map[string]int = map[string]int{}

func NewStartupWindow(a fyne.App, regionStore store.RegionStore, leagueStore store.LeagueStore) fyne.Window {
	w := a.NewWindow("Bygfoot Football Manager")

	widgetStartLeague := &widget.Select{}

	widgetChooseCountry := widget.NewSelect(getCountries(regionStore), func(name string) {
		widgetStartLeague.Options = getLeaguesByCountry(leagueStore, name)
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
	countries, err := store.GetCountries()
	if err != nil {
		log.WithFields(log.Fields{
			"Error": err,
		}).Error("failed to load countries")
	}

	for _, c := range countries {
		countryNameToId[c.Name] = c.Id
	}

	countryNames := make([]string, 0, len(countryNameToId))
	for k := range countryNameToId {
		countryNames = append(countryNames, k)
	}

	return countryNames
}

func getLeaguesByCountry(store store.LeagueStore, countryName string) []string {
	countryId := countryNameToId[countryName]

	leagues, err := store.GetLeaguesByCountryId(countryId)
	if err != nil {
		log.WithFields(log.Fields{
			"Error": err,
		}).Error("failed to load leagues")
	}

	leagueNames := make([]string, 0, len(leagues))
	for _, l := range leagues {
		leagueNames = append(leagueNames, l.Name)
	}

	return leagueNames
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
