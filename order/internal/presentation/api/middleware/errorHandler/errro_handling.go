package errorhandler

import (
	"fmt"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/logger"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/presentation/api/controllers/response"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/presentation/api/middleware/interfaces"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrorMiddleware struct {
	interfaces.Middleware
	logger.Logger
}
type ErrorCatching struct {
	status    *int
	err       error
	exception *response.ExceptionResponse
}
type ErrorStatus struct {
	status    int
	exception any
}

func NewErrorMiddleware(logger logger.Logger) ErrorMiddleware {
	return ErrorMiddleware{
		Logger: logger,
	}
}

func (m ErrorMiddleware) Handle(c *gin.Context) {
	c.Next()
	for _, err := range c.Errors {
		status, exceptionResponse := http.StatusInternalServerError, response.ExceptionResponse{
			Message: "Unknown server error has occurred",
			Data:    err.Error(),
		}
		errorCatching := ErrorCatching{status: &status, err: err, exception: &exceptionResponse}

		handleProductError(errorCatching)
		handleOrderError(errorCatching)
		m.Logger.Info(
			fmt.Sprintf("Server handle erorr with status: %d, and error message: %s",
				*errorCatching.status, errorCatching.err.Error()),
		)
		c.JSON(*errorCatching.status, *errorCatching.exception)
	}
}
