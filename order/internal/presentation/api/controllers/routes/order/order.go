package order

import (
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/presentation/api/controllers/handlers/order"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/presentation/api/engine"
)

type Routes struct {
	engine.GroupRoutes
	controller order.Handler
}

func (r Routes) Setup() {
	r.POST("/orders", r.controller.CreateOrder)
	r.DELETE("/orders", r.controller.DeleteOrder)

	r.GET("/orders", r.controller.GetAllOrders)
	r.GET("/orders/:id", r.controller.GetOrderByID)
	r.GET("/orders/user/:userID", r.controller.GetAllOrdersByUserID)
}

func NewOrderRoutes(
	group engine.GroupRoutes,
	controller order.Handler,
) Routes {
	return Routes{controller: controller, GroupRoutes: group}
}
