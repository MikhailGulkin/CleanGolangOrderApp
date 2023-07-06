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

func NewErrorMiddleware(handler engine.RequestHandler) ErrorMiddleware {
	return ErrorMiddleware{
		RequestHandler: handler,
	}
}

func (m ErrorMiddleware) Handle(c *gin.Context) {
	c.Next()
	err := c.Errors.Last()
	if err != nil {
		if handleError(err, c) {
			return
		}
		c.JSON(http.StatusInternalServerError,
			response.ExceptionResponse{Message: "Unknown server errorHandler has occurred", Data: err.Error()},
		)
	}
}
