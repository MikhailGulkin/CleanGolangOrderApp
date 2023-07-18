package errorhandler

import (
	"errors"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/user/exceptions"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/presentation/api/controllers/response"
	"net/http"
)

func handleUserError(e ErrorCatching) {
	var userIDNotExist *exceptions.UserIDNotExist
	var userAddressIDError *exceptions.UserAddressIDNotExist

	if errors.As(e.err, &userIDNotExist) {
		*e.status = http.StatusNotFound
		response.SetExceptionPayload(e.exception, userIDNotExist.CustomException)
	}
	if errors.As(e.err, &userAddressIDError) {
		*e.status = http.StatusNotFound
		response.SetExceptionPayload(e.exception, userAddressIDError.CustomException)
	}
}
