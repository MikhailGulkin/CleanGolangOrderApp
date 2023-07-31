package brokerconfigurate

import (
	messagebroker "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/messageBroker"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/messageBroker/brokerConfigurate/interfaces"
	"github.com/rabbitmq/amqp091-go"
)

type MessageBrokerConfigure struct {
	interfaces.BaseMessageBrokerConfigure
	Channel *amqp091.Channel
}

func (m *MessageBrokerConfigure) DeclareExchange(exchangeName string) {
	err := m.Channel.ExchangeDeclare(exchangeName, "topic", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}
}
func (m *MessageBrokerConfigure) DeclareQueue(queueName string) {
	_, err := m.Channel.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		panic(err)
	}
}
func (m *MessageBrokerConfigure) BindExchangeQueue(exchangeName, key, queueName string) {
	err := m.Channel.QueueBind(queueName, key, exchangeName, false, nil)
	if err != nil {
		panic(err)
	}
}
func NewMessageBrokerConfigure(rabbit messagebroker.Rabbit) interfaces.BaseMessageBrokerConfigure {
	return &MessageBrokerConfigure{
		Channel: rabbit.GetChannel(),
	}
}
