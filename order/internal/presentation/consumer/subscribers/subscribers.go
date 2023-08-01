package subscribers

import (
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/order/interfaces/cache"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/order/interfaces/saga"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/logger"
	messagebroker "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/messageBroker"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/presentation/consumer/subscribers/order"
)

func NewEventConsumer(
	rabbit messagebroker.Rabbit,
	sagaOrderCreate saga.CreateOrder,
	cache cache.OrderCache,
	logger logger.Logger,
) Subscribers {
	return Subscribers{
		order.SagaCreateSubscriber{ReusableChannel: rabbit.GetChannel(), CreateOrder: sagaOrderCreate, Logger: logger},
		order.OrderEvent{ReusableChannel: rabbit.GetChannel(), Logger: logger, OrderCache: cache},
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
