package errorhandler

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/controllers/response"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/engine"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/middleware/interfaces"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrorMiddleware struct {
	interfaces.Middleware
	engine.RequestHandler
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

func NewErrorMiddleware(handler engine.RequestHandler) ErrorMiddleware {
	return ErrorMiddleware{
		RequestHandler: handler,
	}
}

func (m ErrorMiddleware) Handle(c *gin.Context) {
	c.Next()
	err := c.Errors.Last()

	if err != nil {
		status, exceptionResponse := http.StatusInternalServerError, response.ExceptionResponse{
			Message: "Unknown server error has occurred",
			Data:    err.Error(),
		}
		errorCatching := ErrorCatching{status: &status, err: err, exception: &exceptionResponse}

		handleProductError(errorCatching)
		handleAddressError(errorCatching)
		if c.Request.Method == "POST" ||
			c.Request.Method == "DELETE" ||
			c.Request.Method == "PATCH" ||
			c.Request.Method == "PUT" {
			c.Status(*errorCatching.status)
		}
		if c.Request.Method == "GET" {
			c.JSON(*errorCatching.status, *errorCatching.exception)
		}
	}
}
