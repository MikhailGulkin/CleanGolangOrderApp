package routes

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/config"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/engine"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/middleware"
	"github.com/gin-gonic/gin"
)

type GroupRoutes struct {
	*gin.RouterGroup
}

func NewGroupRoutes(config config.APIConfig, handler engine.RequestHandler, middlewares middleware.Middlewares) GroupRoutes {
	return GroupRoutes{handler.Gin.Group(config.BaseURLPrefix, middlewares...)}
}
