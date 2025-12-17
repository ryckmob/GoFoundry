package app

import (
	"fmt"

	"github.com/ryckmob/GoFoundry/internal/common"
)

func routeTemplate(name string) string {
	return fmt.Sprintf(`package %s

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(r fiber.Router) {
	r.Get("/%s", Get%s)
}
`, name, name, common.Capitalize(name))
}
