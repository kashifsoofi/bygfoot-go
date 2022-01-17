package main

import (
	"os"

	"github.com/kashifsoofi/bygfoot-go/ui"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	log.Debug("main")

	app := ui.NewApp(nil)
	app.Run()
}
