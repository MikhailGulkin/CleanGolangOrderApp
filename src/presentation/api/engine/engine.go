package engine

import (
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

func NewBaseGroup(handler RequestHandler) BaseGroup {
	return BaseGroup{handler.Gin.Group("/api/v1")}
}

var Module = fx.Options(
	fx.Provide(NewRequestHandler),
	fx.Provide(NewBaseGroup),
)
