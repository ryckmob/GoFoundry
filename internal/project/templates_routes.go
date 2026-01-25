package project

func mainRoutesTemplate(project string) string {
	return `package http

// ================================= AVISO =================================
// Arquivo responsável por registrar todas as rotas dos módulos do projeto.
// Importe aqui os apps que possuem rotas a serem expostas e registre-os
// no método RegisterAll.
// ================================= AVISO =================================

import (
	"github.com/gofiber/fiber/v2"
	// "` + project + `/internal/apps/<app-example>"
	// "` + project + `/internal/apps/<app-example2>"
	// "` + project + `/internal/apps/<app-example3>"
)

func RegisterAll(app *fiber.App) {
	// Para cada app importado, chame RegisterRoutes passando o app do Fiber.
	// Exemplo:
	// <app-example>.RegisterRoutes(app)
	// <app-example2>.RegisterRoutes(app)
	// <app-example3>.RegisterRoutes(app)
}
`
}
