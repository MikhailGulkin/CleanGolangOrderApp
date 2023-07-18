package mediator

import "github.com/MikhailGulkin/simpleGoOrderApp/internal/application/user/interfaces/command"

type CreateUserCommandHandler struct {
	command.CreateUser
}

func (c *CreateUserCommandHandler) Handle(cmd interface{}) error {
	return c.Create(cmd.(command.CreateUserCommand))
}
