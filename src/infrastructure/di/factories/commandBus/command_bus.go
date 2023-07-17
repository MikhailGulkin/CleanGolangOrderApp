package commandbus

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/product/interfaces/command"
	c "github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/commandBus"
	"go.uber.org/fx"
	"reflect"
)

func NewCommandBus(createProduct command.CreateProduct) c.CommandBus {
	var commandBus c.CommandBusImpl
	commandBus.CommandHandlers = make(map[reflect.Type]c.CommandHandler)
	commandBus.RegisterCommandHandler(command.CreateProductCommand{}, &c.CreateProductCommandHandler{CreateProduct: createProduct})
	return &commandBus
}

var Module = fx.Provide(
	NewCommandBus,
)
