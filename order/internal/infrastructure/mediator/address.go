package mediator

import "github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/address/interfaces/command"

type CreateAddressCommandHandler struct {
	command.CreateAddress
}

func (c *CreateAddressCommandHandler) Handle(cmd interface{}) error {
	return c.Create(cmd.(command.CreateAddressCommand))
}
