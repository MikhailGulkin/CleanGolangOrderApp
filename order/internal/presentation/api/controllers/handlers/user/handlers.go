package user

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/mediator"
)

func NewUserHandler(
	mediator mediator.Mediator,
) Handler {
	return Handler{
		mediator: mediator,
	}
}
