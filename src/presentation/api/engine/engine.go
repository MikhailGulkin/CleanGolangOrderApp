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

var Module = fx.Options(
	//routes.Module,
	fx.Provide(NewRequestHandler),
	//fx.Provide(NewConfig),
)
