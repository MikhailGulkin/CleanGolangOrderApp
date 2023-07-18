package product

import (
	c "github.com/MikhailGulkin/simpleGoOrderApp/internal/presentation/api/config"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/presentation/api/controllers/handlers/product"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/presentation/api/engine"
)

type Routes struct {
	engine.GroupRoutes
	controller product.Handler
	c.APIConfig
}

func (r Routes) Setup() {
	r.POST("/products", r.controller.CreateProduct)
	r.PUT("/products/:productID/productName", r.controller.UpdateProductName)
	r.GET("/products", r.controller.GetAllProducts)
	r.GET("/products/:productName", r.controller.GetProductByName)
}

func NewProductRoutes(
	group engine.GroupRoutes,
	controller product.Handler,
	config c.APIConfig,
) Routes {
	return Routes{controller: controller, GroupRoutes: group, APIConfig: config}
}
