package product

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/presentation/api/controllers/handlers/product"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/presentation/api/engine"
)

type Routes struct {
	engine.GroupRoutes
	controller product.Handler
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
) Routes {
	return Routes{controller: controller, GroupRoutes: group}
}
