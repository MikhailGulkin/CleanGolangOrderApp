package errorhandler

import (
	"errors"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/order/exceptions"
	domain "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/domain/order/exceptions"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/presentation/api/controllers/response"
	"net/http"
)

func handleOrderError(e ErrorCatching) {
	var orderProductsError *domain.OrderProductsEmpty
	var orderIDError *exceptions.OrderIDNotExist
	var orderProductsIDsError *exceptions.ProductIDsNotExist

	if errors.As(e.err, &orderProductsError) {
		*e.status = http.StatusBadRequest
		response.SetExceptionPayload(e.exception, orderProductsError.CustomException)
	}
	if errors.As(e.err, &orderIDError) {
		*e.status = http.StatusNotFound
		response.SetExceptionPayload(e.exception, orderIDError.CustomException)
	}
	if errors.As(e.err, &orderProductsIDsError) {
		*e.status = http.StatusNotFound
		response.SetExceptionPayload(e.exception, orderProductsIDsError.CustomException)
	}
}
