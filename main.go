package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ryckmob/GoFoundry/internal/app"
	"github.com/ryckmob/GoFoundry/internal/project"
)

func main() {
	if len(os.Args) < 3 {
		usage()
		return
	}

	command := strings.ToLower(os.Args[1])
	name := askLowercaseName(os.Args[2])

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

func askLowercaseName(initial string) string {
	reader := bufio.NewReader(os.Stdin)
	name := initial

	for {
		if name == strings.ToLower(name) {
			return name
		}

		fmt.Println("erro: use apenas letras minúsculas")
		fmt.Print("digite novamente: ")

		input, _ := reader.ReadString('\n')
		name = strings.TrimSpace(input)
	}
}

func usage() {
	fmt.Println("uso:")
	fmt.Println("  gofoundry new <nome-do-projeto>")
	fmt.Println("  gofoundry app <nome-do-app>")
}
