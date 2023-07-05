package middleware

import (
	"errors"
	domain "github.com/MikhailGulkin/simpleGoOrderApp/src/domain/entities"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/engine"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

type ErrorMiddleware struct {
	engine.RequestHandler
}

func NewErrorMiddleware(handler engine.RequestHandler) ErrorMiddleware {
	return ErrorMiddleware{
		handler,
	}
}

func (m ErrorMiddleware) Setup() {
	m.Gin.Use(func(c *gin.Context) {
		c.Next()
		err := c.Errors.Last()
		if err != nil {
			var discountError *domain.InvalidDiscountProductCreation
			var priceError *domain.InvalidPriceProductCreation
			if errors.As(err, &discountError) {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
					"data":  reflect.TypeOf(domain.InvalidDiscountProductCreation{}).String(),
				})
				return
			}
			if errors.As(err, &priceError) {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": err.Error(),
					"data":    reflect.TypeOf(domain.InvalidPriceProductCreation{}).String(),
				})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Unknown server error has occurred",
				"data":    err.Error(),
			})
		}
	})
}
