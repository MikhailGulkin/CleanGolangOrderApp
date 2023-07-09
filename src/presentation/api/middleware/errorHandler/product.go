package errorhandler

import (
	"errors"
	application "github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/exceptions"
	domain "github.com/MikhailGulkin/simpleGoOrderApp/src/domain/entities"
	"net/http"
)

func handleProductError(e ErrorCatching) {
	var discountError *domain.InvalidDiscountProductCreation
	var priceError *domain.InvalidPriceProductCreation
	var productNameError *application.ProductNameNotExist
	var productIDError *application.ProductIDNotExist

	if errors.As(e.err, &discountError) {
		*e.status = http.StatusBadRequest
		e.exception.Message = discountError.Message
		e.exception.Data = discountError.Ctx
	}
	if errors.As(e.err, &priceError) {
		*e.status = http.StatusBadRequest
		e.exception.Message = priceError.Message
		e.exception.Data = priceError.Ctx
	}
	if errors.As(e.err, &productNameError) {
		*e.status = http.StatusNotFound
		e.exception.Message = productNameError.Message
		e.exception.Data = productNameError.Ctx
	}
	if errors.As(e.err, &productIDError) {
		*e.status = http.StatusNotFound
		e.exception.Message = productIDError.Message
		e.exception.Data = productIDError.Ctx
	}
}
