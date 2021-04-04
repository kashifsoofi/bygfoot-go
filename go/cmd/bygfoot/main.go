package main

import (
	"github.com/kashifsoofi/bygfoot-go/db"
	log "github.com/sirupsen/logrus"
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
}
