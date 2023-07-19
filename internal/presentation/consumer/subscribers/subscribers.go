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
	eventSub := events.NewOrderEventQuery(logger, cache, connection)
	return Subscribers{
		order.SagaCreateSubscriber{Channel: ch, CreateOrder: sagaOrderCreate, Logger: logger},
		eventSub,
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
