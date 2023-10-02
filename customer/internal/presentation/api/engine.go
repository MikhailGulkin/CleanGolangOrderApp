package api

import (
	"github.com/MikhailGulkin/CleanGolangOrderApp/pkg/env"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

type Engine struct {
	app *fiber.App
}

func NewEngine() Engine {
	return Engine{app: fiber.New()}
}
func (e *Engine) Run() error {
	return e.app.Listen(env.GetEnv("API_PORT", ":8000"))
}
func (e *Engine) Setup() {
	e.app.Get("/swagger/*", swagger.HandlerDefault)
}
func (e *Engine) Shutdown() error {
	return e.app.Shutdown()
}
