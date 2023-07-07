package errorhandler

import (
	"errors"
	application "github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/exceptions"
	domain "github.com/MikhailGulkin/simpleGoOrderApp/src/domain/entities"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/controllers/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func handleError(err error, c *gin.Context) bool {
	var discountError *domain.InvalidDiscountProductCreation
	var priceError *domain.InvalidPriceProductCreation
	var productNameError *application.ProductNameNotExist
	var productIDError *application.ProductIDNotExist

	if errors.As(err, &discountError) {
		c.JSON(http.StatusBadRequest,
			response.ExceptionResponse{Message: discountError.Message, Data: discountError.Ctx},
		)
		return true
	}
	if errors.As(err, &priceError) {
		c.JSON(http.StatusBadRequest,
			response.ExceptionResponse{Message: priceError.Message, Data: priceError.Ctx},
		)
		return true
	}
	if errors.As(err, &productNameError) {
		c.JSON(http.StatusNotFound,
			response.ExceptionResponse{Message: productNameError.Message, Data: productNameError.Ctx},
		)
		return true
	}
	if errors.As(err, &productIDError) {
		c.JSON(http.StatusNotFound,
			response.ExceptionResponse{Message: productIDError.Message, Data: productIDError.Ctx},
		)
		return true
	}
	return false
}
