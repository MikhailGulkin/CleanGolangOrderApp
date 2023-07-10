package errorhandler

import (
	"errors"
	application "github.com/MikhailGulkin/simpleGoOrderApp/src/application/product/exceptions"
	domain "github.com/MikhailGulkin/simpleGoOrderApp/src/domain/aggregate/product"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/controllers/response"
	"net/http"
)

type ErrorStatus struct {
	status    int
	exception any
}

func handleProductError(e ErrorCatching) {
	var discountError *domain.InvalidDiscountProductCreation
	var priceError *domain.InvalidPriceProductCreation
	var incorrectProductName *domain.InvalidProductNameUpdate
	var productNameError *application.ProductNameNotExist
	var productIDError *application.ProductIDNotExist

	if errors.As(e.err, &discountError) {
		*e.status = http.StatusBadRequest
		response.SetExceptionPayload(e.exception, discountError.CustomException)
	}
	if errors.As(e.err, &priceError) {
		*e.status = http.StatusBadRequest
		response.SetExceptionPayload(e.exception, priceError.CustomException)
	}
	if errors.As(e.err, &productNameError) {
		*e.status = http.StatusNotFound
		response.SetExceptionPayload(e.exception, productNameError.CustomException)
	}
	if errors.As(e.err, &productIDError) {
		*e.status = http.StatusNotFound
		response.SetExceptionPayload(e.exception, productIDError.CustomException)
	}
	if errors.As(e.err, &incorrectProductName) {
		*e.status = http.StatusBadRequest
		response.SetExceptionPayload(e.exception, incorrectProductName.CustomException)
	}
}
