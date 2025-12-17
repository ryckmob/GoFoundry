package project

func mainTemplate(project string) string {
	return `package main

import (
	"github.com/gofiber/fiber/v2"

	httpRoutes "` + project + `/internal/http"
)

func main() {
	app := fiber.New()

	httpRoutes.RegisterAll(app)

	app.Listen(":8000")
}
`
}
