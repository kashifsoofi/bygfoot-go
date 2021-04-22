package main

import (
	"github.com/kashifsoofi/bygfoot-go/store"
	"github.com/kashifsoofi/bygfoot-go/ui"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	store := store.NewSqlStore("sqlite3", "bygfoot.db")
	defer store.Close()

	app := ui.NewApp(store)
	app.Run()
}
