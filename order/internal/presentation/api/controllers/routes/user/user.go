package user

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/presentation/api/controllers/handlers/user"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/presentation/api/engine"
)

type Routes struct {
	engine.GroupRoutes
	controller user.Handler
}

func (r Routes) Setup() {
	r.POST("/users", r.controller.CreateUser)
}

func NewUserRoutes(
	group engine.GroupRoutes,
	controller user.Handler,
) Routes {
	return Routes{controller: controller, GroupRoutes: group}
}
