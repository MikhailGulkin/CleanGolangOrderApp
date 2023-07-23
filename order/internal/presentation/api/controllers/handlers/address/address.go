package address

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/mediator"
)

func NewProductHandler(
	mediator mediator.Mediator,
) Handler {
	return Handler{
		mediator: mediator,
	}
}
