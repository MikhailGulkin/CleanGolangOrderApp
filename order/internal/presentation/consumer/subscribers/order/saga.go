package order

import (
	"encoding/json"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/order/interfaces/saga"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/logger"
	messagebroker "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/messageBroker"
)

type SagaCreateSubscriber struct {
	messagebroker.Rabbit
	logger.Logger
	saga.CreateOrder
}

func (s SagaCreateSubscriber) Listen() {
	messages, _ := s.Rabbit.GetChannel().Consume(
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
				s.Info("Error unmarshal event type; err %s", err.Error())
			}
			s.CreateOrder.CheckStatus(m)
		}
	}()
}
