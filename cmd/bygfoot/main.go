package main

import (
	"os"

	"log/slog"

	"github.com/jessevdk/go-flags"
	"github.com/kashifsoofi/bygfoot-go/internal/store/file"
	"github.com/kashifsoofi/bygfoot-go/internal/ui/gtk4"
	"go.uber.org/zap"
	"go.uber.org/zap/exp/zapslog"
	"go.uber.org/zap/zapcore"
)

type Options struct {
	// -1 = Debug
	LogLevel zapcore.Level `short:"l" long:"loglevel" description:"Set current log level" default:"-1"`
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

	config := zap.Config{
		Level:            zap.NewAtomicLevelAt(options.LogLevel),
		Development:      true,
		Encoding:         "json",
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	zapLogger := zap.Must(config.Build())
	defer zapLogger.Sync()

	logger := slog.New(zapslog.NewHandler(zapLogger.Core(), nil))
	slog.SetDefault(logger)
}

func main() {
	slog.Debug("main")

	// app := fyne.NewApp(nil)
	store := file.NewFileStore()
	app := gtk4.NewApp(store)
	app.Run()
}
