package address

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/presentation/api/controllers/handlers/address"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/presentation/api/engine"
)

type Routes struct {
	engine.GroupRoutes
	controller address.Handler
}

func (r Routes) Setup() {
	r.POST("/addresses", r.controller.CreateAddress)
}

func NewAddressRoutes(
	group engine.GroupRoutes,
	controller address.Handler,
) Routes {
	return Routes{controller: controller, GroupRoutes: group}
}
