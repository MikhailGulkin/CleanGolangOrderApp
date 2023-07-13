package errorhandler

import (
	"errors"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/exceptions"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/controllers/response"
	"net/http"
)

func handleOrderError(e ErrorCatching) {
	var orderProductsError *exceptions.OrderProductsEmpty

	if errors.As(e.err, &orderProductsError) {
		*e.status = http.StatusBadRequest
		response.SetExceptionPayload(e.exception, orderProductsError.CustomException)
	}
}
