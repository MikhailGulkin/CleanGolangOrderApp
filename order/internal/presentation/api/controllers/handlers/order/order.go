package order

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/mediator"
)

func NewOrderHandler(
	mediator mediator.Mediator,
) Handler {
	return Handler{
		mediator: mediator,
	}
}
