package file

import (
	"bufio"
	"errors"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"log/slog"

	"github.com/kashifsoofi/bygfoot-go/internal/domain"
	"github.com/kashifsoofi/bygfoot-go/internal/store"
)

const homeDirName = ".bygfoot"

type FileStore struct {
	hints store.HintsStore
}

func NewFileStore() store.Store {
	return FileStore{
		hints: newHintsStore(),
	}
}

func (s FileStore) Close() {}

func (s FileStore) Hints() store.HintsStore {
	return s.hints
}

func (s FileStore) League() store.LeagueStore { return nil }

func (s FileStore) Region() store.RegionStore { return nil }

func exists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

func getHomeDir() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		slog.Error("cannot get user home dir")
		panic(err)
	}
	return dirname
}

func getCurrentDir() string {
	ex, err := os.Executable()
	if err != nil {
		slog.Error("cannot get current path")
		panic(err)
	}

	currentDir := filepath.Dir(ex)
	return currentDir
}

func osIsUnix() bool {
	return runtime.GOOS != "windows"
}

func getBygfootDir() string {
	home := getHomeDir()

	if osIsUnix() {
		return path.Join(home, homeDirName)
	}

	return getCurrentDir()
}

func findSupportFile(filename string) string {
	return ""
}

func loadOptionsFile(filename string) *domain.OptionList {
	//support_file := findSupportFile(filename)
	support_file := filename
	file, err := os.Open(support_file)
	if err != nil {
		slog.Error("cannot open file", slog.String("support_file", support_file))
		panic(err)
	}
	defer file.Close()

	optionList := &domain.OptionList{}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		optStr := scanner.Text()
		name, value, found := strings.Cut(optStr, " ")
		if !found {
			continue
		}

		option := domain.Option{Name: name}
		if strings.HasPrefix(name, "string_") {
			option.StringValue, option.Value = value, -1
		} else {
			option.Value, _ = strconv.Atoi(value)
		}
		optionList.Add(option)

		newName := ""
		if osIsUnix() && strings.HasSuffix(name, "_unix") {
			newName, _ = strings.CutSuffix(name, "_unix")
		} else if !osIsUnix() && strings.HasSuffix(name, "_win32") {
			newName, _ = strings.CutSuffix(name, "_win32")
		}
		if newName != "" {
			newOption := domain.Option{Name: newName, StringValue: value}
			optionList.Add(newOption)
		}
	}

	return optionList
}
