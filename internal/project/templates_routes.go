package project

func mainRoutesTemplate(project string) string {
	return `package http

import (
	"github.com/gofiber/fiber/v2"
	"` + project + `/internal/apps/app-example"
)

func RegisterAll(app *fiber.App) {
	// Nenhuma rota registrada ainda
	// app-example.RegisterRoutes(app)
}
`
}
