package main

import (
	"log"

	"fyne.io/fyne/v2/app"
	"github.com/kashifsoofi/bygfoot-go/db"
	"github.com/kashifsoofi/bygfoot-go/ui"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	sqliteDb, err := db.NewDB("bygfoot.db")
	if err != nil {
		log.Fatal(err)
	}
	defer sqliteDb.Close()

	err = db.RunMigrations(sqliteDb)
	if err != nil {
		log.Fatal(err)
	}

	a := app.New()
	w := ui.NewSplashWindow(a)

	w.ShowAndRun()
}
