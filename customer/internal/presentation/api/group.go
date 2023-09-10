package api

import "github.com/gofiber/fiber/v2"

type FiberGroup struct {
	fiber.Router
}

func NewFiberGroup(engine Engine, middlewares Middlewares) FiberGroup {
	return FiberGroup{Router: engine.Group("/api/v1", middlewares...)}
}
