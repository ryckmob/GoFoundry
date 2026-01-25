package app

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ryckmob/GoFoundry/internal/common"
)

type Field struct {
	Name string
	Type string
}

var allowedTypes = map[string]bool{
	"string":    true,
	"int":       true,
	"int64":     true,
	"int32":     true,
	"float32":   true,
	"float64":   true,
	"bool":      true,
	"byte":      true,
	"rune":      true,
	"time.Time": true,
}

func findAppsDir(baseDir string) (string, error) {
	var appsDir string

	err := filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && info.Name() == "apps" {
			appsDir = path
			return filepath.SkipDir
		}
		return nil
	})

	if err != nil {
		return "", err
	}

	if appsDir == "" {
		appsDir = filepath.Join("internal", "apps")
		if err := os.MkdirAll(appsDir, 0755); err != nil {
			return "", err
		}
	}

	return appsDir, nil
}

func parseFields(args []string) ([]Field, error) {
	var fields []Field
	seen := make(map[string]bool)

	for _, arg := range args {
		parts := strings.Split(arg, ":")
		if len(parts) != 2 {
			return nil, fmt.Errorf("campo inválido: %s", arg)
		}

		rawName := strings.TrimSpace(parts[0])
		typ := strings.TrimSpace(parts[1])

		if rawName == "" {
			return nil, fmt.Errorf("nome de campo vazio")
		}

		if !allowedTypes[typ] {
			return nil, fmt.Errorf("tipo não suportado: %s", typ)
		}

		name := common.Capitalize(rawName)

		if seen[name] {
			return nil, fmt.Errorf("campo duplicado: %s", rawName)
		}
		seen[name] = true

		fields = append(fields, Field{
			Name: name,
			Type: typ,
		})
	}

	return fields, nil
}

// Gera o SQL da tabela baseado nos fields
func schemaSQL(name string, fields []Field) string {
	tableName := strings.ToLower(name) // tabela em minúsculas
	var lines []string

	// Cria tabela formatada
	lines = append(lines, "CREATE TABLE "+tableName+" (")
	lines = append(lines, "  id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,")

	for i, f := range fields {
		colName := strings.ToLower(f.Name)
		var colType string
		switch f.Type {
		case "string":
			colType = "VARCHAR(255) NOT NULL"
		case "int", "int32", "int64":
			colType = "INT NOT NULL"
		case "float32", "float64":
			colType = "FLOAT NOT NULL"
		case "bool":
			colType = "TINYINT(1) NOT NULL"
		case "byte", "rune":
			colType = "VARCHAR(10) NOT NULL"
		case "time.Time":
			colType = "DATETIME NOT NULL"
		default:
			colType = "VARCHAR(255) NOT NULL"
		}

		line := fmt.Sprintf("  %s %s", colName, colType)
		if i < len(fields)-1 {
			line += ","
		}
		lines = append(lines, line)
	}

	lines = append(lines, ");")
	lines = append(lines, "")

	// Cria comando único em uma linha e comenta de forma elegante
	singleLine := "CREATE TABLE " + tableName + " (id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY"
	for _, f := range fields {
		colName := strings.ToLower(f.Name)
		var colType string
		switch f.Type {
		case "string":
			colType = "VARCHAR(255) NOT NULL"
		case "int", "int32", "int64":
			colType = "INT NOT NULL"
		case "float32", "float64":
			colType = "FLOAT NOT NULL"
		case "bool":
			colType = "TINYINT(1) NOT NULL"
		case "byte", "rune":
			colType = "VARCHAR(10) NOT NULL"
		case "time.Time":
			colType = "DATETIME NOT NULL"
		default:
			colType = "VARCHAR(255) NOT NULL"
		}
		singleLine += ", " + colName + " " + colType
	}
	singleLine += ");"

	lines = append(lines, "-- Para criar esta tabela em uma única execução, utilize o comando abaixo:")
	lines = append(lines, "-- "+singleLine)

	return strings.Join(lines, "\n")
}

func CreateApp(name string, rawFields []string) error {
	appsBase, err := findAppsDir(".")
	if err != nil {
		return err
	}

	base := filepath.Join(appsBase, name)
	if err := os.MkdirAll(base, 0755); err != nil {
		return err
	}

	var fields []Field
	if len(rawFields) > 0 {
		fields, err = parseFields(rawFields)
		if err != nil {
			return err
		}
	}

	// Cria models.go
	if err := common.CreateFile(base+"/models.go", modelTemplate(name, fields)); err != nil {
		return err
	}

	// Cria handlers.go
	code, err := handlerTemplate(name, fields)
	if err != nil {
		return err
	}
	if err := common.CreateFile(base+"/handlers.go", code); err != nil {
		return err
	}

	// Cria routes.go
	if err := common.CreateFile(base+"/routes.go", routeTemplate(name)); err != nil {
		return err
	}

	// Cria schema.sql
	sql := schemaSQL(name, fields)
	if err := common.CreateFile(base+"/schema.sql", sql); err != nil {
		return err
	}

	fmt.Println("App criado:", name, "em", appsBase)
	return nil
}
