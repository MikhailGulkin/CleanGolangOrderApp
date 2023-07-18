package commanddispatcher

import (
	"reflect"
)

type CommandDispatcher interface {
	Send(command interface{}) error
	RegisterCommandHandler(command interface{}, handler CommandHandler)
}

type CommandDispatcherImpl struct {
	CommandHandlers map[reflect.Type]CommandHandler
}
type CommandHandler interface {
	Handle(command interface{}) error
}

func (c *CommandDispatcherImpl) Send(command interface{}) error {
	var handler CommandHandler
	c.GetCommandHandler(command, &handler)
	return handler.Handle(command)
}
func (c *CommandDispatcherImpl) GetCommandHandler(command interface{}, receiver interface{}) {
	reflect.ValueOf(receiver).Elem().Set(reflect.ValueOf(c.CommandHandlers[reflect.TypeOf(command)]))
}

func (c *CommandDispatcherImpl) RegisterCommandHandler(command interface{}, handler CommandHandler) {
	c.CommandHandlers[reflect.ValueOf(command).Type()] = handler
}
