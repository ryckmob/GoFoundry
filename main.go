package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ryckmob/GoFoundry/internal/app"
	"github.com/ryckmob/GoFoundry/internal/project"
	projectfromfile "github.com/ryckmob/GoFoundry/internal/project_from_file"
)

const (
	reset = "\033[0m"
	green = "\033[32m"
	blue  = "\033[34m"
	cyan  = "\033[36m"
	bold  = "\033[1m"
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
		if err != nil {
			break
		}

		fmt.Println()
		fmt.Println(bold + cyan + "Projeto criado com sucesso" + reset)
		fmt.Println("Nome:", bold+name+reset)
		fmt.Println()
		fmt.Println(blue + "PASSO 1:" + reset)
		fmt.Println(blue + "Entre no diretório do projeto para começar" + reset)
		fmt.Println()
		fmt.Println(blue + "PASSO 2:" + reset)
		fmt.Println(
			"gofoundry app <nome-do-app> --struct " +
				"nome:" + green + "string" + reset + " " +
				"idade:" + green + "int" + reset + " " +
				"ativo:" + green + "bool" + reset,
		)
		fmt.Println()
		fmt.Println(bold + "Tipos de dados suportados:" + reset)
		fmt.Println(green + "int      int64      float32      float64" + reset)
		fmt.Println(green + "string   bool       byte         int32      time.Time" + reset)
		fmt.Println()

	case "app":
		fields := parseStructArgs(os.Args)
		err = app.CreateApp(name, fields)

	case "startproject":
		if len(os.Args) < 3 {
			fmt.Println("erro: informe o arquivo .gofoundry")
			return
		}
		err = projectfromfile.Run(os.Args[2])

	default:
		usage()
		return
	}

	if err != nil {
		fmt.Println("erro:", err)
		os.Exit(1)
	}
}

func parseStructArgs(args []string) []string {
	for i, arg := range args {
		if arg == "--struct" && i+1 < len(args) {
			return args[i+1:]
		}
	}
	return nil
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
	fmt.Println()
	fmt.Println(bold + cyan + "GoFoundry CLI" + reset)
	fmt.Println(blue + "gerador de projetos e apps em Go" + reset)
	fmt.Println()

	fmt.Println(bold + "uso:" + reset)
	fmt.Println()

	fmt.Println(green + "  gofoundry new <nome-do-projeto>" + reset)
	fmt.Println("    cria a estrutura base de um novo projeto")
	fmt.Println()

	fmt.Println(green + "  gofoundry app <nome-do-app>" + reset)
	fmt.Println("    cria um app vazio dentro do projeto atual")
	fmt.Println()

	fmt.Println(green + "  gofoundry app <nome-do-app> --struct nome:string idade:int ativo:bool" + reset)
	fmt.Println("    cria um app já com a struct definida")
	fmt.Println()

	fmt.Println(green + "  gofoundry startproject <arquivo.gofoundry>" + reset)
	fmt.Println("    cria um projeto completo a partir de um único arquivo declarativo")
	fmt.Println()

	fmt.Println(cyan + "dica:" + reset)
	fmt.Println("  o comando startproject cria toda a estrutura do sistema de forma automática;")
	fmt.Println("  gera models, rotas e CRUDs prontos a partir de um único arquivo.")
	fmt.Println("  veja um exemplo do arquivo .gofoundry no repositório do GitHub")
	fmt.Println()
}
