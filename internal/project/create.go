package project

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/ryckmob/GoFoundry/environment"
	"github.com/ryckmob/GoFoundry/internal/common"
)

func CreateProject(name string) error {
	base := name

	dirs := []string{
		base + "/cmd",
		base + "/internal",
		base + "/internal/apps",
		base + "/internal/http",
		base + "/database",
	}

	for _, d := range dirs {
		if err := os.MkdirAll(d, 0755); err != nil {
			return err
		}
	}

	if err := common.CreateFile(base+"/cmd/main.go", mainTemplate(name)); err != nil {
		return err
	}
	if err := common.CreateFile(base+"/internal/http/routes.go", mainRoutesTemplate(name)); err != nil {
		return err
	}

	if err := common.CreateFile(base+"/database/connection.go", databaseConnectionTemplate()); err != nil {
		return err
	}

	err := environment.CreateEnv(name)
	if err != nil {
		fmt.Println("erro ao criar .env:", err)
		os.Exit(1)
	}

	cmd := exec.Command("go", "mod", "init", name)
	cmd.Dir = base
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("erro ao rodar go mod init: %v\n%s", err, string(output))
	}

	cmd = exec.Command("go", "mod", "tidy")
	cmd.Dir = base
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("erro ao executar go mod tidy: %v\n%s", err, string(output))
	}

	fmt.Println("Dependências instaladas com sucesso.")
	fmt.Println("Projeto inicializado e pronto para desenvolvimento.")
	return nil
}
