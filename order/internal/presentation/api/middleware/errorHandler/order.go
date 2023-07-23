package errorhandler

import (
	"errors"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/order/exceptions"
	domain "github.com/MikhailGulkin/simpleGoOrderApp/order/internal/domain/order/exceptions"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/presentation/api/controllers/response"
	"net/http"
)

func handleOrderError(e ErrorCatching) {
	var orderProductsError *domain.OrderProductsEmpty
	var orderIDError *exceptions.OrderIDNotExist
	var orderClientIDError *exceptions.OrderClientIDNotExist
	var orderProductsIDsError *exceptions.ProductIDsNotExist
	var orderAddressIDError *exceptions.OrderAddressIDNotExist

	if errors.As(e.err, &orderProductsError) {
		*e.status = http.StatusBadRequest
		response.SetExceptionPayload(e.exception, orderProductsError.CustomException)
	}
	if errors.As(e.err, &orderIDError) {
		*e.status = http.StatusNotFound
		response.SetExceptionPayload(e.exception, orderIDError.CustomException)
	}
	if errors.As(e.err, &orderClientIDError) {
		*e.status = http.StatusNotFound
		response.SetExceptionPayload(e.exception, orderClientIDError.CustomException)
	}
	if errors.As(e.err, &orderProductsIDsError) {
		*e.status = http.StatusNotFound
		response.SetExceptionPayload(e.exception, orderProductsIDsError.CustomException)
	}
	if errors.As(e.err, &orderAddressIDError) {
		*e.status = http.StatusNotFound
		response.SetExceptionPayload(e.exception, orderAddressIDError.CustomException)
	}
}
