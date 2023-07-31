package product

import (
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/mediator"
)

func NewProductHandler(
	mediator mediator.Mediator,
) Handler {
	return Handler{
		mediator: mediator,
	}
}
