package main

import (
	"os"

	"github.com/SteaceP/optimum_vape/http/route"
	_ "github.com/SteaceP/optimum_vape/http/validation"

	"goyave.dev/goyave/v4"
	_ "goyave.dev/goyave/v4/database/dialect/postgres"
)

func main() {
	// This is the entry point of your application.
	if err := goyave.Start(route.Register); err != nil {
		os.Exit(err.(*goyave.Error).ExitCode)
	}
}
