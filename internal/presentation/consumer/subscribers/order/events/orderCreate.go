package events

import (
	"encoding/json"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/interfaces/cache"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/logger"
	"github.com/rabbitmq/amqp091-go"
	"sync"
)

type CreateQuerySubscriber struct {
	*amqp091.Channel
	logger.Logger
	cache.OrderCache
}

func (s CreateQuerySubscriber) Listen(mutex sync.Locker) {
	messages, _ := s.Channel.Consume(
		"Orders",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	var e cache.OrderCreateSubscribe
	var str string
	go func() {
		for message := range messages {
			_ = json.Unmarshal(message.Body, &str)
			err := json.Unmarshal([]byte(str), &e)
			if err != nil {
				continue
			}
			mutex.Lock()
			s.OrderCache.OrderEvent(e)
			mutex.Unlock()
		}
	}()
}
