package middleware

import (
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/presentation/api/middleware/errorHandler"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/presentation/api/middleware/logging"
	"github.com/gin-gonic/gin"
)

type Middlewares []gin.HandlerFunc

func NewMiddlewares(
	errorMiddleware errorhandler.ErrorMiddleware,
	loggingMiddleware logging.LoggerMiddleware,
) Middlewares {
	return Middlewares{
		errorMiddleware.Handle,
		loggingMiddleware.Handle,
	}
}
