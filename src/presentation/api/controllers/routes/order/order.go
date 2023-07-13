package order

import (
	c "github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/config"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/controllers/handlers/order"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/engine"
)

type Routes struct {
	engine.GroupRoutes
	controller order.Handler
	c.APIConfig
}

func (r Routes) Setup() {
	r.POST("/orders", r.controller.CreateOrder)
}

func NewOrderRoutes(
	group engine.GroupRoutes,
	controller order.Handler,
	config c.APIConfig,
) Routes {
	return Routes{controller: controller, GroupRoutes: group, APIConfig: config}
}
