package main

import (
	"time"

	"github.com/tak1827/zerolog-playground/log"
)

func main() {
	log.Logger.Debug().Msg("hello world")

	// set log level
	log.SetLevel(log.INFO_LEVEL)
	log.Logger.Debug().Msg("hello world, again") // no output
	log.Logger.Info().Msg("this is warning!")

	// set writor
	log.SetWriter(log.ConsoleWriter)
	log.Logger.Info().Msg("pretty logging")

	// set time format
	log.SetTimeFormat(time.RFC3339)

	// moduleA
	loggerA := log.ModuleA("event")
	loggerA.Info().Msg("something")

	// moduleB
	loggerB := log.ModuleB("")
	loggerB.Info().Msg("something")
}
