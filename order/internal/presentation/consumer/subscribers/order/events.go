package order

import (
	"encoding/json"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/order/interfaces/cache"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/logger"
	"github.com/MikhailGulkin/CleanGolangOrderApp/pkg/rabbit"
)

type OrderEvent struct {
	*rabbit.ReusableChannel
	logger.Logger
	cache.OrderCache
}

func (s OrderEvent) Listen() {
	defer s.ReusableChannel.Release()
	// TODO: ADD ACK FOR EVENTS
	messages, _ := s.ReusableChannel.Consume(
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
	var orderDelete cache.OrderDeleteEvent
	var eventType cache.EventType
	var str string
	go func() {
		for message := range messages {
			_ = json.Unmarshal(message.Body, &str)
			err := json.Unmarshal([]byte(str), &eventType)
			if err != nil {
				s.Info("Error unmarshal event type; err %s", err.Error())
			}
			switch eventType.EventType {
			case "OrderCreated":
				err = json.Unmarshal([]byte(str), &orderCreate)
				if err != nil {
					continue
				}
				s.OrderCache.OrderCreateEvent(orderCreate)
			case "OrderAddProduct":
				err = json.Unmarshal([]byte(str), &orderAddProduct)
				if err != nil {
					continue
				}
				s.OrderCache.OrderAddProductEvent(orderAddProduct)
			case "OrderDeleted":
				err = json.Unmarshal([]byte(str), &orderDelete)
				if err != nil {
					continue
				}
				s.OrderCache.OrderDeleteEvent(orderDelete)
			}
		}
	}()
}
