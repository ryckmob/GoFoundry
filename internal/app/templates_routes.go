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

// RegisterRoutes registra as rotas do Fiber
func RegisterRoutes(r fiber.Router) {
	r.Post("/%s", Insert%sHandler)
}
`, name, name, common.Capitalize(name))
}
