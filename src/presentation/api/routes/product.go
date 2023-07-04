package routes

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/controllers"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/engine"
)

type ProductRoutes struct {
	handler    engine.RequestHandler
	controller controllers.ProductController
}

func (s ProductRoutes) Setup() {
	api := s.handler.Gin.Group("/api")
	{
		api.POST("/product", s.controller.CreateProduct)
		//api.GET("/user", s.userController.GetUser)
		//api.GET("/user/:id", s.userController.GetOneUser)
		//api.POST("/user", s.userController.SaveUser)
		//api.POST("/user/:id", s.userController.UpdateUser)
		//api.DELETE("/user/:id", s.userController.DeleteUser)
	}
}

func NewProductRoutes(
	handler engine.RequestHandler,
	controller controllers.ProductController,
) ProductRoutes {
	return ProductRoutes{handler: handler, controller: controller}
}
