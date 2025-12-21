package app

import (
	"fmt"

	"github.com/ryckmob/GoFoundry/internal/common"
)

func handlerTemplate(name string) (string, error) {
	project, err := FindProjectRoot(".")
	if err != nil {
		return "", err
	}

	importPath := project + "/database"
	typeName := common.Capitalize(name)

	return fmt.Sprintf(`package %s

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"%s"
)

// Insert%s insere no banco de dados
func Insert%s(%s %sModel) (sql.Result, error) {
	return database.DB.Exec(
		"INSERT INTO %ss (nome, preco) VALUES (?, ?)",
		%s.Teste1,
		%s.Teste2,
	)
}

// Insert%sHandler é o handler do Fiber
func Insert%sHandler(c *fiber.Ctx) error {
	var %s %sModel
	if err := c.BodyParser(&%s); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
	}

	_, err := Insert%s(%s)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to insert %s"})
	}

	return c.Status(201).JSON(fiber.Map{"message": "%s inserido com sucesso"})
}
`,
		// argumentos fmt.Sprintf
		name,       // package %s
		importPath, // import "%s"

		typeName, // Insert%s
		typeName, // Insert%s(%s %sModel)
		name,     // (%s %sModel)
		typeName, // %sModel
		name,     // %ss
		name,     // %s.Teste1
		name,     // %s.Teste2

		typeName, // Insert%sHandler
		typeName, // Insert%sHandler
		name,     // var %s
		typeName, // %sModel
		name,     // &%s BodyParser
		typeName, // Insert%s
		name,     // (%s)
		name,     // error message
		name,     // success message
	), nil
}
