package response

import (
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/domain/common"
)

type ExceptionResponse struct {
	Message string `json:"message"`
	Data    string `json:"data"`
}

func SetExceptionPayload(response *ExceptionResponse, exc common.CustomException) {
	response.Message = exc.Message
	response.Data = exc.Ctx
}
