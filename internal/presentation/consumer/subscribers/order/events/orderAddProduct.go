package events

import (
	"encoding/json"
	"fmt"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/interfaces/cache"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/logger"
	"github.com/rabbitmq/amqp091-go"
	"sync"
)

type AddProductQuerySubscriber struct {
	*amqp091.Channel
	logger.Logger
	cache.OrderCache
}

func (s AddProductQuerySubscriber) Listen(mutex sync.Locker) {
	err := s.Channel.QueueBind(
		"OrdersAddProduct",
		"Order.AddProduct",
		"Orders",
		false,
		nil,
	)
	if err != nil {
		return
	}

	messages, _ := s.Channel.Consume(
		"OrdersAddProduct",
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	var e cache.OrderAddProductSubscribe
	var str string
	go func() {
		for message := range messages {
			_ = json.Unmarshal(message.Body, &str)
			err := json.Unmarshal([]byte(str), &e)
			if err != nil {
				s.Info(fmt.Sprintf("Invalid order unmarshall id %s, err %s", message.MessageId, err.Error()))
				continue
			}
			mutex.Lock()
			s.OrderCache.OrderCreate(e)
			mutex.Unlock()
		}
	}()
}
