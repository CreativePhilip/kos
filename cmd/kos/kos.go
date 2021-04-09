package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	setupCli()
}

func setupCli() {
	app := &cli.App{
		Name:  "KurumaOS CLI",
		Usage: "Set of utilities for developing distributed applications with the KurumaOS framework",
		Commands: []*cli.Command{
			GetInitProjectCliOption(),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
