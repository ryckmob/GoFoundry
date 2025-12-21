package project

func mainTemplate(project string) string {
	return `package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"` + project + `/database"
	httpRoutes "` + project + `/internal/http"
)

func main() {
	logFile, err := os.OpenFile("startup.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return
	}
	defer logFile.Close()

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(logFile),
		zap.InfoLevel,
	)

	logger := zap.New(core)
	defer logger.Sync()

	app := fiber.New()

	if err := database.Connect(); err != nil {
		logger.Error("erro ao conectar no banco", zap.Error(err))
		return
	}

	defer database.DB.Close()

	httpRoutes.RegisterAll(app)

	if err := app.Listen(":8000"); err != nil {
		logger.Error("erro ao subir servidor", zap.Error(err))
	}
}
`
}
