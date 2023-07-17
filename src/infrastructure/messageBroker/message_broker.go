package messagebroker

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/messageBroker/config"
	"github.com/rabbitmq/amqp091-go"
)

func BuildDial(config config.MessageBrokerConfig) *amqp091.Connection {
	conn, err := amqp091.Dial(config.FullDNS())
	if err != nil {
		panic(err)
	}
	return conn
}
func BuildChannel(connection *amqp091.Connection) *amqp091.Channel {
	ch, err := connection.Channel()
	if err != nil {
		panic(err)
	}
	return ch
}
