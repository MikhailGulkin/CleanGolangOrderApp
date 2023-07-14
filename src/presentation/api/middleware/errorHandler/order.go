package errorhandler

import (
	"errors"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/exceptions"
	domain "github.com/MikhailGulkin/simpleGoOrderApp/src/domain/order/exceptions"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/controllers/response"
	"net/http"
)

func handleOrderError(e ErrorCatching) {
	var orderProductsError *domain.OrderProductsEmpty
	var orderIDError *exceptions.OrderIDNotExist
	if errors.As(e.err, &orderProductsError) {
		*e.status = http.StatusBadRequest
		response.SetExceptionPayload(e.exception, orderProductsError.CustomException)
	}
	if errors.As(e.err, &orderIDError) {
		*e.status = http.StatusNotFound
		response.SetExceptionPayload(e.exception, orderIDError.CustomException)
	}
}