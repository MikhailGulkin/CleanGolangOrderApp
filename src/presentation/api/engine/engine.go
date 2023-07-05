package engine

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type RequestHandler struct {
	Gin *gin.Engine
}

func NewRequestHandler() RequestHandler {
	engine := gin.Default()
	return RequestHandler{Gin: engine}
}

type BaseGroup struct {
	Group *gin.RouterGroup
}

func NewBaseGroup(handler RequestHandler, config config.APIConfig) BaseGroup {
	return BaseGroup{handler.Gin.Group(config.BaseURLPrefix)}
}

var Module = fx.Options(
	fx.Provide(NewRequestHandler),
	fx.Provide(NewBaseGroup),
)
