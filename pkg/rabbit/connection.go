package rabbit

import "github.com/rabbitmq/amqp091-go"

// connect establish the connection with the broker amqp
func connect(connectionString string) (*amqp091.Connection, error) {
	connection, err := amqp091.Dial(connectionString)
	if err != nil {
		return nil, err
	}

	return connection, nil
}
