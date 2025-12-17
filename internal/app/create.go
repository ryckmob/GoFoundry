package app

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ryckmob/GoFoundry/internal/common"
)

// Procura a primeira pasta chamada "apps" no diretório atual e subdiretórios
func findAppsDir(baseDir string) (string, error) {
	var appsDir string
	err := filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && info.Name() == "apps" {
			appsDir = path
			return filepath.SkipDir // já achou, pode parar
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	if appsDir == "" {
		// se não encontrou, usa o padrão
		appsDir = filepath.Join("internal", "apps")
		if err := os.MkdirAll(appsDir, 0755); err != nil {
			return "", err
		}
	}
	return appsDir, nil
}

func CreateApp(name string) error {
	appsBase, err := findAppsDir(".")
	if err != nil {
		return err
	}

	base := filepath.Join(appsBase, name)
	if err := os.MkdirAll(base, 0755); err != nil {
		return err
	}

	if err := common.CreateFile(base+"/models.go", modelTemplate(name)); err != nil {
		return err
	}
	if err := common.CreateFile(base+"/handlers.go", handlerTemplate(name)); err != nil {
		return err
	}
	if err := common.CreateFile(base+"/routes.go", routeTemplate(name)); err != nil {
		return err
	}
	if err := common.CreateFile(base+"/service.go", serviceTemplate(name)); err != nil {
		return err
	}

	fmt.Println("App criado:", name, "em", appsBase)
	return nil
}
