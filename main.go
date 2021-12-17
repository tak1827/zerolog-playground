package main

import (
	"os"
	"time"

	"github.com/tak1827/zerolog-playground/log"
)

func main() {
	log.Logger.Debug().Msg("hello world")

	// set log level
	log.SetLevel(log.INFO_LEVEL)
	log.Logger.Debug().Msg("hello world, again") // no output
	log.Logger.Info().Msg("this is info")

	// set writor
	log.SetWriter(os.Stderr)
	log.Logger.Info().Msg("os.Stderr")

	// set time format
	log.SetTimeFormat(time.Kitchen)

	// moduleA
	loggerA := log.ModuleA("event")
	loggerA.Info().Msg("something")

	// moduleB
	loggerB := log.ModuleB("")
	loggerB.Info().Msg("something")
}
