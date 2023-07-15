package controller

import (
	"context"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/common/interfaces/broker"
	"github.com/rabbitmq/amqp091-go"
)

type MessageBrokerImpl struct {
	broker.MessageBroker
	*amqp091.Channel
}

func (m *MessageBrokerImpl) PublishMessage(ctx context.Context, exchangeName, routingKey string, message []byte) error {
	err := m.PublishWithContext(
		ctx,
		exchangeName,
		routingKey,
		true,
		false,
		m.BuildMessage(message),
	)
	return err
}
func (m *MessageBrokerImpl) BuildMessage(message []byte) amqp091.Publishing {
	return amqp091.Publishing{
		ContentType: "application/json",
		Body:        message,
	}
}
func NewMessageBroker(channel *amqp091.Channel) broker.MessageBroker {
	return &MessageBrokerImpl{
		Channel: channel,
	}
}
