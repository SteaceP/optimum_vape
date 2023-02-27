package main

import (
	"github.com/SteaceP/optimum_vape/http/route"
	_ "github.com/SteaceP/optimum_vape/http/validation"
	"os"

	goyave "goyave.dev/goyave/v4"
	_ "goyave.dev/goyave/v4/database/dialect/postgres"
)

func main() {
	if err := goyave.Start(route.Register); err != nil {
		os.Exit(err.(*goyave.Error).ExitCode)

	}

}
