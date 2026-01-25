package environment

import (
	"os"
)

// CreateEnv cria um arquivo .env padrão dentro da pasta do projeto
// Retorna erro caso não consiga criar o arquivo
func CreateEnv(projectName string) error {
	envPath := projectName + "/.env"
	envContent := `DB_USER=root
DB_PASS=
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=meubanco
`

	err := os.WriteFile(envPath, []byte(envContent), 0644)
	if err != nil {
		return err
	}

	return nil
}
