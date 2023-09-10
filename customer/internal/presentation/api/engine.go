package api

import "github.com/gofiber/fiber/v2"

type Engine struct {
	*fiber.App
}

func NewEngine() Engine {
	return Engine{App: fiber.New()}
}
