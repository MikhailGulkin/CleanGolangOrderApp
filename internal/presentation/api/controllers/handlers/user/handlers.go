package user

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/mediator"
)

func NewUserHandler(
	mediator mediator.Mediator,
) Handler {
	return Handler{
		mediator: mediator,
	}
}
