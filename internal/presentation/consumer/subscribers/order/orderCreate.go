package order

import (
	"encoding/json"
	"fmt"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/dto"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/logger"
	"github.com/rabbitmq/amqp091-go"
)

type CreateQuerySubscriber struct {
	*amqp091.Channel
	logger.Logger
}

func (s CreateQuerySubscriber) Listen() {
	err := s.Channel.QueueBind(
		"OrdersCreate",
		"order_create",
		"Orders",
		false,
		nil,
	)
	if err != nil {
		return
	}

	messages, _ := s.Channel.Consume(
		"OrdersCreate",
		"order_create",
		false,
		false,
		false,
		false,
		nil,
	)

	var e dto.OrderCreateSubscribe
	var str string
	go func() {
		for message := range messages {
			_ = json.Unmarshal(message.Body, &str)
			err := json.Unmarshal([]byte(str), &e)
			if err != nil {
				s.Info(fmt.Sprintf("Invalid order unmarshall id %s, err %s", message.MessageId, err.Error()))
			}
		}
	}()
}
