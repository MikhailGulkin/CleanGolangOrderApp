package address

import (
	c "github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/config"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/controllers/handlers/address"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/engine"
)

type Routes struct {
	engine.GroupRoutes
	controller address.Handler
	c.APIConfig
}

func (r Routes) Setup() {
	r.POST("/address", r.controller.CreateAddress)
}

func NewAddressRoutes(
	group engine.GroupRoutes,
	controller address.Handler,
	config c.APIConfig,
) Routes {
	return Routes{controller: controller, GroupRoutes: group, APIConfig: config}
}
