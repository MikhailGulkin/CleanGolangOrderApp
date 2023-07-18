package order

import (
	c "github.com/MikhailGulkin/simpleGoOrderApp/internal/presentation/api/config"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/presentation/api/controllers/handlers/order"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/presentation/api/engine"
)

type Routes struct {
	engine.GroupRoutes
	controller order.Handler
	c.APIConfig
}

func (r Routes) Setup() {
	r.POST("/orders", r.controller.CreateOrder)
	r.GET("/orders", r.controller.GetAllOrders)
	r.GET("/orders/:id", r.controller.GetOrderByID)
}

func NewOrderRoutes(
	group engine.GroupRoutes,
	controller order.Handler,
	config c.APIConfig,
) Routes {
	return Routes{controller: controller, GroupRoutes: group, APIConfig: config}
}
