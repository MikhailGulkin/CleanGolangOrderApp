package subscribers

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/interfaces/saga"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/logger"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/presentation/consumer/subscribers/order"
	"github.com/rabbitmq/amqp091-go"
)

func NewEventConsumer(
	connection *amqp091.Connection,
	sagaOrderCreate saga.CreateOrder,
	logger logger.Logger,
) Subscribers {
	ch, err := connection.Channel()
	ch1, err1 := connection.Channel()
	if err != nil || err1 != nil {
		panic("Incorrect event consumer create")
	}
	return Subscribers{
		order.SagaCreateSubscriber{Channel: ch, CreateOrder: sagaOrderCreate, Logger: logger},
		order.CreateQuerySubscriber{Channel: ch1, Logger: logger},
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
