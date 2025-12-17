package main

import (
	"fmt"
	"os"

	"github.com/ryckmob/GoFoundry/internal/app"
	"github.com/ryckmob/GoFoundry/internal/project"
)

func main() {
	if len(os.Args) < 3 {
		usage()
		return
	}

	command := os.Args[1]
	name := os.Args[2]

	var err error

	switch command {
	case "new":
		err = project.CreateProject(name)
	case "app":
		err = app.CreateApp(name)
	default:
		usage()
		return
	}

	if err != nil {
		fmt.Println("erro:", err)
		os.Exit(1)
	}
}

func usage() {
	fmt.Println("uso:")
	fmt.Println("  gofoundry new <nome-do-projeto>")
	fmt.Println("  gofoundry app <nome-do-app>")
}
