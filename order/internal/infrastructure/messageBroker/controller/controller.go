package controller

import (
	"context"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/common/interfaces/broker"
	messagebroker "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/messageBroker"
	"github.com/MikhailGulkin/CleanGolangOrderApp/pkg/rabbit"
	"github.com/rabbitmq/amqp091-go"
)

type MessageBrokerImpl struct {
	broker.MessageBroker
	*rabbit.ReusableChannel
}

func (m *MessageBrokerImpl) PublishMessage(ctx context.Context, exchangeName, routingKey string, message []byte) error {
	err := m.ReusableChannel.PublishWithContext(
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
func NewMessageBroker(rabbit messagebroker.Rabbit) broker.MessageBroker {
	return &MessageBrokerImpl{
		ReusableChannel: rabbit.GetChannel(),
	}
}
