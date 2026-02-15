package projectfromfile

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ryckmob/GoFoundry/internal/app"
	"github.com/ryckmob/GoFoundry/internal/project"
)

func Run(filePath string) error {
	if !strings.HasSuffix(filePath, ".gofoundry") {
		return fmt.Errorf("arquivo inválido: apenas arquivos com extensão .gofoundry são permitidos")
	}

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var hasAnyLine bool
	var hasProject bool

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		hasAnyLine = true

		if strings.HasPrefix(line, "Project ") {
			projectName := strings.TrimSpace(strings.TrimPrefix(line, "Project"))

			if projectName != strings.ToLower(projectName) {
				return fmt.Errorf("nome do projeto inválido: use apenas letras minúsculas")
			}

			err := project.CreateProject(projectName)
			if err != nil {
				return err
			}

			err = os.Chdir(projectName)
			if err != nil {
				return fmt.Errorf("erro ao entrar no diretório do projeto: %v", err)
			}

			hasProject = true
			continue
		}

		if strings.HasPrefix(line, "app ") {
			if !hasProject {
				return fmt.Errorf("erro de definição: a linha Project deve vir antes dos apps")
			}

			err := runAppLine(line)
			if err != nil {
				return err
			}
			continue
		}

		return fmt.Errorf("linha inválida no arquivo: %s", line)
	}

	if !hasAnyLine {
		return fmt.Errorf("arquivo inválido: arquivo vazio")
	}

	if !hasProject {
		return fmt.Errorf("arquivo inválido: é obrigatório declarar a linha Project")
	}

	return scanner.Err()
}

func runAppLine(line string) error {
	parts := strings.Fields(line)

	if len(parts) < 2 {
		return fmt.Errorf("definição de app inválida: %s", line)
	}

	appName := parts[1]

	if appName != strings.ToLower(appName) {
		return fmt.Errorf("nome do app inválido: use apenas letras minúsculas")
	}

	var structFields []string
	for i, p := range parts {
		if p == "--struct" && i+1 < len(parts) {
			structFields = parts[i+1:]
			break
		}
	}

	return app.CreateApp(appName, structFields)
}
