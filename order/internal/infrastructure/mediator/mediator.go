package mediator

import (
	c "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/mediator/dispatchers/commandDispatcher"
	q "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/mediator/dispatchers/queryDispatcher"
	"reflect"
)

type Mediator interface {
	Send(command interface{}) (interface{}, error)
	Query(query interface{}) (interface{}, error)
	RegisterCommandHandler(command interface{}, handler c.CommandHandler)
	RegisterQueryHandler(command interface{}, handler q.QueryHandler)
}

type MediatorImpl struct {
	queryDispatcher   q.QueryDispatcher
	commandDispatcher c.CommandDispatcher
}

func (mediator *MediatorImpl) RegisterCommandHandler(command interface{}, handler c.CommandHandler) {
	mediator.commandDispatcher.RegisterCommandHandler(command, handler)
}

func (mediator *MediatorImpl) RegisterQueryHandler(command interface{}, handler q.QueryHandler) {
	mediator.queryDispatcher.RegisterQueryHandler(command, handler)
}

func (mediator *MediatorImpl) Send(command interface{}) (interface{}, error) {
	return mediator.commandDispatcher.Send(command)
}

func (mediator *MediatorImpl) Query(query interface{}) (interface{}, error) {
	return mediator.queryDispatcher.Query(query)
}

func (MediatorImpl) Create() Mediator {
	var commandDispatcher c.CommandDispatcherImpl
	var queryDispatcher q.QueryDispatcherImpl
	queryDispatcher.QueryHandlers = make(map[reflect.Type]q.QueryHandler)
	commandDispatcher.CommandHandlers = make(map[reflect.Type]c.CommandHandler)
	return &MediatorImpl{queryDispatcher: &queryDispatcher, commandDispatcher: &commandDispatcher}
}
