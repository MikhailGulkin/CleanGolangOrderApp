package routes

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/controllers/handlers"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/engine"
	"github.com/gin-gonic/gin"
)

type ProductRoutes struct {
	engine.BaseGroup
	controller handlers.ProductHandler
}

func (s ProductRoutes) Setup() {
	s.Group.POST("/products", s.controller.CreateProduct)
	s.Group.GET("hello", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"hello": "world",
		})
	})
}

func NewProductRoutes(
	group engine.BaseGroup,
	controller handlers.ProductHandler,
) ProductRoutes {
	return ProductRoutes{controller: controller, BaseGroup: group}
}
