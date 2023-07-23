package logging

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/logger"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/presentation/api/middleware/interfaces"
	"github.com/gin-gonic/gin"
)

type LoggerMiddleware struct {
	logger.Logger
	interfaces.Middleware
}

func NewLoggerMiddleware(logger logger.Logger) LoggerMiddleware {
	return LoggerMiddleware{
		Logger: logger,
	}
}
func (m LoggerMiddleware) Handle(c *gin.Context) {
	c.Next()
	m.Logger.Infow(
		"Api good answer Request",
		"Request type", c.Request.Method,
		"Request status", c.Writer.Status(),
		"Content length", c.Writer.Size(),
		"Request path", c.Request.URL,
	)
}
