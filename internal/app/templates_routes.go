package app

import (
	"fmt"

	"github.com/ryckmob/GoFoundry/internal/common"
)

func routeTemplate(name string) string {
	typeName := common.Capitalize(name)

	return fmt.Sprintf(`package %s

import (
	"github.com/gofiber/fiber/v2"
)

// RegisterRoutes registra as rotas do Fiber
func RegisterRoutes(r fiber.Router) {
	r.Post("/%s", Insert%sHandler)
	r.Put("/%s", Update%sHandler)
	r.Delete("/%s/:id", Delete%sHandler)
	r.Get("/%s", List%ssHandler)
}
`,
		name,

		name, typeName,
		name, typeName,
		name, typeName,
		name, typeName,
	)
}
