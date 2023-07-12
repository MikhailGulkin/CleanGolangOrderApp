package engine

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type RequestHandler struct {
	Gin *gin.Engine
}
type GroupRoutes struct {
	*gin.RouterGroup
}

func NewRequestHandler() RequestHandler {
	engine := gin.Default()
	return RequestHandler{Gin: engine}
}

var Module = fx.Options(
	fx.Provide(NewRequestHandler),
)
