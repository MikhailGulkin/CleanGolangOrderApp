package order

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/mediator"
)

func NewOrderHandler(
	mediator mediator.Mediator,
) Handler {
	return Handler{
		mediator: mediator,
	}
}
