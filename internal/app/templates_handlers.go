package app

import (
	"fmt"

	"github.com/ryckmob/GoFoundry/internal/common"
)

func handlerTemplate(name string) string {
	return fmt.Sprintf(`package %s

import (
	"github.com/gofiber/fiber/v2"
)

func Get%s(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"app": "%s",
	})
}
`, name, common.Capitalize(name), name)
}
