package commandbus

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/product/interfaces/command"
	"reflect"
)

type CommandBus interface {
	Execute(command interface{}) error
}

type CommandBusImpl struct {
	CommandHandlers map[reflect.Type]CommandHandler
}
type CommandHandler interface {
	handle(command interface{}) error
}

func (c *CommandBusImpl) Execute(command interface{}) error {
	var handler CommandHandler
	c.GetCommandHandler(command, &handler)
	return handler.handle(command)
}
func (c *CommandBusImpl) GetCommandHandler(command interface{}, receiver interface{}) {
	reflect.ValueOf(receiver).Elem().Set(reflect.ValueOf(c.CommandHandlers[reflect.TypeOf(command)]))
}

type CreateProductCommandHandler struct {
	command.CreateProduct
}

func (c *CommandBusImpl) RegisterCommandHandler(command interface{}, handler CommandHandler) {
	c.CommandHandlers[reflect.ValueOf(command).Type()] = handler
}
func (c *CreateProductCommandHandler) handle(cmd interface{}) error {
	return c.Create(cmd.(command.CreateProductCommand))
}
