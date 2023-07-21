package events

import (
	"encoding/json"
	"fmt"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/interfaces/cache"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/logger"
	"github.com/rabbitmq/amqp091-go"
)

type OrderEvent struct {
	*amqp091.Channel
	logger.Logger
	cache.OrderCache
}

func (s OrderEvent) Listen() {
	messages, _ := s.Channel.Consume(
		"Orders",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	var orderAddProduct cache.OrderAddProductSubscribe
	var orderCreate cache.OrderCreateSubscribe
	var eventType cache.EventType
	var str string
	go func() {
		for message := range messages {
			_ = json.Unmarshal(message.Body, &str)
			err := json.Unmarshal([]byte(str), &eventType)
			if err != nil {
				continue
			}
			switch eventType.EventType {
			case "OrderCreated":
				err = json.Unmarshal([]byte(str), &orderCreate)
				if err != nil {
					continue
				}
				fmt.Println("orderCreate")
				s.OrderCache.OrderEvent(orderCreate)
			case "OrderAddProduct":
				err = json.Unmarshal([]byte(str), &orderAddProduct)
				if err != nil {
					continue
				}
				fmt.Println("orderAddProduct")
				s.OrderCache.OrderEvent(orderAddProduct)
			}
		}
	}()
}
