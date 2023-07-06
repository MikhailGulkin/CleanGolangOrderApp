package middleware

import (
	"errors"
	domain "github.com/MikhailGulkin/simpleGoOrderApp/src/domain/entities"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/controllers/response"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/engine"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrorMiddleware struct {
	Middleware
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
		var discountError *domain.InvalidDiscountProductCreation
		var priceError *domain.InvalidPriceProductCreation
		if errors.As(err, &discountError) {
			c.JSON(http.StatusBadRequest,
				response.ExceptionResponse{Message: discountError.Message, Data: discountError.Ctx},
			)
			return
		}
		if errors.As(err, &priceError) {
			c.JSON(http.StatusBadRequest,
				response.ExceptionResponse{Message: priceError.Message, Data: priceError.Ctx},
			)
			return
		}

		c.JSON(http.StatusInternalServerError,
			response.ExceptionResponse{Message: "Unknown server error has occurred", Data: err.Error()},
		)
	}
}
