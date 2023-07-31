package rabbit

import "github.com/rabbitmq/amqp091-go"

func newReusableChannel(connection *amqp091.Connection) (*amqp091.Channel, error) {
	channel, err := connection.Channel()
	if err != nil {
		return nil, err
	}

	return channel, nil
}
