package main

import (
	"os"

	"github.com/kashifsoofi/bygfoot-go/ui"
	log "github.com/sirupsen/logrus"

	"github.com/jessevdk/go-flags"
)

type Options struct {
	// 2 = Error, 5 = Debug
	LogLevel log.Level `short:"l" long:"loglevel" description:"Set current log level" default:"5"`
}

var options Options
var parser = flags.NewParser(&options, flags.Default)

func init() {
	if _, err := parser.Parse(); err != nil {
		switch flagsErr := err.(type) {
		case flags.ErrorType:
			if flagsErr == flags.ErrHelp {
				os.Exit(0)
			}
			os.Exit(1)
		default:
			os.Exit(1)
		}
	}

	log.SetOutput(os.Stdout)
	log.SetLevel(options.LogLevel)
}

func main() {
	log.Debug("main debug")

	app := ui.NewApp(nil)
	app.Run()
}
