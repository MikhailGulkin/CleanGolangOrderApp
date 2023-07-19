package middleware

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/presentation/api/middleware/errorHandler"
	"github.com/gin-gonic/gin"
)

type Middlewares []gin.HandlerFunc

func NewMiddlewares(
	errorMiddleware errorhandler.ErrorMiddleware,
) Middlewares {
	return Middlewares{
		errorMiddleware.Handle,
	}
}