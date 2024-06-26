package main

import (
	"github.com/codelite7/momentum/api/cmd/generate"
	"github.com/codelite7/momentum/api/cmd/run"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			run.RunCommand,
			generate.GenerateCommand,
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
