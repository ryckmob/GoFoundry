package app

import (
	"fmt"
	"strings"

	"github.com/ryckmob/GoFoundry/internal/common"
)

func handlerTemplate(name string, fields []Field) (string, error) {
	project, err := FindProjectRoot(".")
	if err != nil {
		return "", err
	}

	importPath := project + "/database"
	typeName := common.Capitalize(name)
	tableName := name

	var (
		columns     []string
		insertMarks []string
		insertArgs  []string
		updateSets  []string
		updateArgs  []string
		selectCols  []string
		scanArgs    []string
	)

	for _, f := range fields {
		col := strings.ToLower(f.Name)
		columns = append(columns, col)
		insertMarks = append(insertMarks, "?")
		insertArgs = append(insertArgs, "data."+f.Name)
		updateSets = append(updateSets, col+" = ?")
		updateArgs = append(updateArgs, "data."+f.Name)
		selectCols = append(selectCols, col)
		scanArgs = append(scanArgs, "&item."+f.Name)
	}

	return fmt.Sprintf(`package %s

import (
	"github.com/gofiber/fiber/v2"
	"%s"
)

func Insert%sHandler(c *fiber.Ctx) error {
	var data %sModel
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "request invalido"})
	}

	_, err := database.DB.Exec(
		"INSERT INTO %s (%s) VALUES (%s)",
		%s,
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"message": "%s criado com sucesso"})
}

func Update%sHandler(c *fiber.Ctx) error {
	var data %sModel
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "request invalido"})
	}

	_, err := database.DB.Exec(
		"UPDATE %s SET %s WHERE id = ?",
		%s,
		data.ID,
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "%s atualizado com sucesso"})
}

func Delete%sHandler(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "id invalido"})
	}

	_, err = database.DB.Exec(
		"DELETE FROM %s WHERE id = ?",
		id,
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "%s removido com sucesso"})
}

func List%ssHandler(c *fiber.Ctx) error {
	rows, err := database.DB.Query(
		"SELECT id, %s FROM %s",
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	defer rows.Close()

	var items []%sModel
	for rows.Next() {
		var item %sModel
		if err := rows.Scan(&item.ID, %s); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		items = append(items, item)
	}

	return c.JSON(items)
}
`,
		name,
		importPath,

		typeName,
		typeName,
		tableName,
		strings.Join(columns, ", "),
		strings.Join(insertMarks, ", "),
		strings.Join(insertArgs, ", "),
		name,

		typeName,
		typeName,
		tableName,
		strings.Join(updateSets, ", "),
		strings.Join(updateArgs, ", "),
		name,

		typeName,
		tableName,
		name,

		typeName,
		strings.Join(selectCols, ", "),
		tableName,
		typeName,
		typeName,
		strings.Join(scanArgs, ", "),
	), nil
}
