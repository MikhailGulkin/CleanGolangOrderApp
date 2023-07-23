package engine

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/logger"
	api "github.com/MikhailGulkin/simpleGoOrderApp/internal/presentation/api/config"
	"github.com/gin-gonic/gin"
)

type RequestHandler struct {
	Gin *gin.Engine
}
type GroupRoutes struct {
	*gin.RouterGroup
}

func NewRequestHandler(logger logger.Logger, config api.APIConfig) RequestHandler {
	gin.DefaultWriter = logger.GetGinLogger()
	if config.Mode == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	return RequestHandler{Gin: gin.New()}
}
