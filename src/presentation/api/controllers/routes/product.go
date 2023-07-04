package routes

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/controllers/handlers"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/engine"
)

type ProductRoutes struct {
	engine.BaseGroup
	controller handlers.ProductHandler
}

func (s ProductRoutes) Setup() {
	s.Group.POST("/product", s.controller.CreateProduct)
}

func NewProductRoutes(
	group engine.BaseGroup,
	controller handlers.ProductHandler,
) ProductRoutes {
	return ProductRoutes{controller: controller, BaseGroup: group}
}
