package order

import (
	"encoding/json"
	"fmt"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/interfaces/saga"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/logger"
	"github.com/rabbitmq/amqp091-go"
)

type SagaCreateSubscriber struct {
	*amqp091.Channel
	logger.Logger
	saga.CreateOrder
}

func (s SagaCreateSubscriber) Listen() {
	err := s.Channel.QueueBind(
		"OrderSagaStatus",
		"Order.Saga.Status",
		"Orders",
		false,
		nil,
	)
	if err != nil {
		return
	}

	messages, _ := s.Channel.Consume(
		"OrderSagaStatus",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	var m saga.Message
	go func() {
		for message := range messages {
			err := json.Unmarshal(message.Body, &m)
			if err != nil {
				s.Info(fmt.Sprintf("Invalid order unmarshall id %s, err %s", message.MessageId, err.Error()))
				continue
			}
			s.CreateOrder.CheckStatus(m)
		}
	}()
}
