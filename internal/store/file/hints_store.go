package file

import (
	"fmt"
	"os"
	"path"
	"strconv"

	"github.com/kashifsoofi/bygfoot-go/internal/domain"
	"github.com/kashifsoofi/bygfoot-go/internal/store"
)

type hintsStore struct {
	hints *domain.OptionList
}

func newHintsStore() store.HintsStore {
	return hintsStore{
		hints: loadHintsFile(),
	}
}

func (s hintsStore) LoadHintNumber() int {
	filename := getHintsFilename()
	if !exists(filename) {
		return 0
	}

	b, err := os.ReadFile(filename)
	if err != nil {
		return 0
	}

	hintNum, err := strconv.Atoi(string(b))
	if err != nil {
		return 0
	}

	return hintNum
}

func (s hintsStore) SaveHintNumber(hintNum int) {
	filename := getHintsFilename()
	os.WriteFile(filename, []byte(strconv.Itoa(hintNum)), 0644)
}

func (s hintsStore) Hint(hintNum int) string {
	return s.hints.ListItem(hintNum).StringValue
}

func (s hintsStore) TotalHints() int {
	return s.hints.Length() - 1
}

func getHintsFilename() string {
	bygfootDir := getBygfootDir()
	return path.Join(bygfootDir, "hint_num")
}

func getLanguageCode() string {
	return "en"
}

func loadHintsFile() *domain.OptionList {
	hintsFilename := fmt.Sprintf("bygfoot_hints_%s", getLanguageCode())
	hintsFilenameSup := findSupportFile(hintsFilename)
	if hintsFilenameSup == "" {
		hintsFilename = "support_files/hints/bygfoot_hints_en"
	}

	return loadOptionsFile(hintsFilename)
}
