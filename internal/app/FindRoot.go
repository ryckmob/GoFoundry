package app

import (
	"fmt"
	"os"
	"path/filepath"
)

func FindProjectRoot(start string) (string, error) {
	current, err := filepath.Abs(start)
	if err != nil {
		return "", err
	}

	for {
		if existsDir(filepath.Join(current, "cmd")) &&
			existsDir(filepath.Join(current, "internal")) &&
			existsDir(filepath.Join(current, "database")) {

			// retorna apenas o nome da pasta raiz
			return filepath.Base(current), nil
		}

		parent := filepath.Dir(current)
		if parent == current {
			break
		}

		current = parent
	}

	return "", fmt.Errorf("Projeto não encontrado!")
}

func existsDir(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}
