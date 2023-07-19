package events

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/interfaces/cache"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/logger"
	"github.com/rabbitmq/amqp091-go"
	"sync"
)

func NewOrderEventQuery(logger logger.Logger, cache cache.OrderCache, connection *amqp091.Connection) OrderEventQuery {
	ch, err := connection.Channel()
	if err != nil {
		panic(err)
	}
	return OrderEventQuery{
		AddProductQuerySubscriber: AddProductQuerySubscriber{Channel: ch, Logger: logger, OrderCache: cache},
		CreateQuerySubscriber:     CreateQuerySubscriber{Channel: ch, Logger: logger, OrderCache: cache},
	}
}

type OrderEventQuery struct {
	AddProductQuerySubscriber
	CreateQuerySubscriber
}

func (w OrderEventQuery) Listen() {
	var mutex sync.Mutex
	go w.AddProductQuerySubscriber.Listen(&mutex)
	go w.CreateQuerySubscriber.Listen(&mutex)
}
