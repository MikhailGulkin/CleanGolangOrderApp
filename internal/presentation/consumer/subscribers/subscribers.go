package subscribers

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/interfaces/cache"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/interfaces/saga"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/logger"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/presentation/consumer/subscribers/order"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/presentation/consumer/subscribers/order/events"
	"github.com/rabbitmq/amqp091-go"
)

func NewEventConsumer(
	connection *amqp091.Connection,
	sagaOrderCreate saga.CreateOrder,
	cache cache.OrderCache,
	logger logger.Logger,
) Subscribers {
	ch, err := connection.Channel()
	if err != nil {
		panic(err)
	}

	ch2, err2 := connection.Channel()
	if err2 != nil {
		panic(err2)
	}
	return Subscribers{
		order.SagaCreateSubscriber{Channel: ch, CreateOrder: sagaOrderCreate, Logger: logger},
		events.OrderEvent{Channel: ch2, Logger: logger, OrderCache: cache},
	}
}

type Subscriber interface {
	Listen()
}
type Subscribers []Subscriber

func (w Subscribers) Listen() {
	for _, worker := range w {
		go worker.Listen()
	}
}
