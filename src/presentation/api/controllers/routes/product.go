package routes

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/config"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/controllers/handlers"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/engine"
)

type ProductRoutes struct {
	engine.RequestHandler
	controller handlers.ProductHandler
	config.APIConfig
}

func (s ProductRoutes) Setup() {
	r := s.Gin.Group(s.BaseURLPrefix)
	r.POST("/products", s.controller.CreateProduct)
	r.GET("/products", s.controller.GetALlProducts)
}

func NewProductRoutes(
	handler engine.RequestHandler,
	controller handlers.ProductHandler,
	config config.APIConfig,
) ProductRoutes {
	return ProductRoutes{controller: controller, RequestHandler: handler, APIConfig: config}
}
