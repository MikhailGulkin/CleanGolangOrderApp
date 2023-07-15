package brokerconfigurate

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/message_broker/broker_configurate/interfaces"
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
func NewMessageBrokerConfigure(ch *amqp091.Channel) interfaces.BaseMessageBrokerConfigure {
	return &MessageBrokerConfigure{
		Channel: ch,
	}
}
