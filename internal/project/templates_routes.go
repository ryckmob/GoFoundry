package project

func mainRoutesTemplate(project string) string {
	return `package http

import (
	"github.com/gofiber/fiber/v2"
	//"` + project + `/internal/apps/<app-example>"
)

func RegisterAll(app *fiber.App) {
	// Registre aqui as rotas de cada app
	// Substitua <app-example> pelo nome do app criado
	// Exemplo:
	// User.RegisterRoutes(app)
}
`
}
