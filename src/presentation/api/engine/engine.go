package engine

import (
	"github.com/gin-gonic/gin"
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
