package errorhandler

import (
	"errors"
	application "github.com/MikhailGulkin/simpleGoOrderApp/internal/application/address/exceptions"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/presentation/api/controllers/response"
	"net/http"
)

func handleAddressError(e ErrorCatching) {
	var addressIDError *application.AddressIDNotExist

	if errors.As(e.err, &addressIDError) {
		*e.status = http.StatusNotFound
		response.SetExceptionPayload(e.exception, addressIDError.CustomException)
	}
}
