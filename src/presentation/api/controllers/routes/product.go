package routes

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/config"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/controllers/handlers"
)

type ProductRoutes struct {
	GroupRoutes
	controller handlers.ProductHandler
	config.APIConfig
}

func (r ProductRoutes) Setup() {
	r.POST("/products", r.controller.CreateProduct)
	r.GET("/products", r.controller.GetALlProducts)
}

func NewProductRoutes(
	group GroupRoutes,
	controller handlers.ProductHandler,
	config config.APIConfig,
) ProductRoutes {
	return ProductRoutes{controller: controller, GroupRoutes: group, APIConfig: config}
}
