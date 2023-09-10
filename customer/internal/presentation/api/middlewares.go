package api

import "github.com/gofiber/fiber/v2"

type Middlewares []fiber.Handler

func NewMiddlewares() Middlewares {
	return Middlewares{}
}
