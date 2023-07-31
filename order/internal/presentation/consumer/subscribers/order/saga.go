package order

import (
	"encoding/json"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/order/interfaces/saga"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/logger"
	"github.com/rabbitmq/amqp091-go"
)

type SagaCreateSubscriber struct {
	*amqp091.Channel
	logger.Logger
	saga.CreateOrder
}

func (s SagaCreateSubscriber) Listen() {
	messages, _ := s.Channel.Consume(
		"CustomerSaga",
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
				continue
			}
			s.CreateOrder.CheckStatus(m)
		}
	}()
}
