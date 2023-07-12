package routes

import (
	c "github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/config"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/controllers/handlers"
)

type ProductRoutes struct {
	GroupRoutes
	controller handlers.ProductHandler
	c.APIConfig
}

func (r ProductRoutes) Setup() {
	r.POST("/products", r.controller.CreateProduct)
	r.PUT("/products/:productID/productName", r.controller.UpdateProductName)
	r.GET("/products", r.controller.GetALlProducts)
	r.GET("/products/:productName", r.controller.GetProductByName)
}

func NewProductRoutes(
	group GroupRoutes,
	controller handlers.ProductHandler,
	config c.APIConfig,
) ProductRoutes {
	return ProductRoutes{controller: controller, GroupRoutes: group, APIConfig: config}
}
