package order

import (
	"encoding/json"
	"fmt"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/saga"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/logger"
	"github.com/rabbitmq/amqp091-go"
)

type SagaCreateSubscriber struct {
	*amqp091.Channel
	logger.Logger
	saga.CreateOrder
}

func (s SagaCreateSubscriber) Listen() {
	err := s.Channel.QueueBind(
		"OrderSagaCreateStatus",
		"saga_order_creation_status",
		"OrderSagaCreate",
		false,
		nil,
	)
	if err != nil {
		return
	}

	messages, _ := s.Channel.Consume(
		"OrderSagaCreateStatus",
		"saga_order_get",
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
			}
			s.CreateOrder.CheckStatus(m)
		}
	}()
}
