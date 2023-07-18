package user

import (
	c "github.com/MikhailGulkin/simpleGoOrderApp/internal/presentation/api/config"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/presentation/api/controllers/handlers/user"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/presentation/api/engine"
)

type Routes struct {
	engine.GroupRoutes
	controller user.Handler
	c.APIConfig
}

func (r Routes) Setup() {
	r.POST("/users", r.controller.CreateUser)
}

func NewUserRoutes(
	group engine.GroupRoutes,
	controller user.Handler,
	config c.APIConfig,
) Routes {
	return Routes{controller: controller, GroupRoutes: group, APIConfig: config}
}
