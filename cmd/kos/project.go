package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"path"
)

func GetInitProjectCliOption() *cli.Command {
	return &cli.Command{
		Name:   "initialize",
		Usage:  "Initialize an empty KurumaOS project",
		Action: onInitProject,
	}
}

func createProjectSkeleton() {
	dir, _ := os.UserHomeDir()
	projectPath := path.Join(dir, "kos_workspace")
	directories := []string{"build", "launches", "messages", "services", "src", ".kos"}

	for _, directory := range directories {
		fullPath := path.Join(projectPath, directory)
		log.Printf("Created directory %s", fullPath)
		err := os.MkdirAll(fullPath, 0700)

		if err != nil {
			log.Fatal(err.Error())
		}
	}
}

func onInitProject(_ *cli.Context) error {
	fmt.Print("Creating a new project KurumaOS project\n\n")

	createProjectSkeleton()

	fmt.Println("\nDone !! Happy Coding and remember")
	fmt.Println("When in doubt remember 42 is the purpose of life :^)")

	return nil
}
