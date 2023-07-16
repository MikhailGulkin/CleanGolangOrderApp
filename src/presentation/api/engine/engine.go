package engine

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/logger"
	"github.com/gin-gonic/gin"
)

type RequestHandler struct {
	Gin *gin.Engine
}
type GroupRoutes struct {
	*gin.RouterGroup
}

func NewRequestHandler(logger logger.Logger) RequestHandler {
	gin.DefaultWriter = logger.GetGinLogger()
	engine := gin.Default()
	return RequestHandler{Gin: engine}
}
