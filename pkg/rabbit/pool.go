package rabbit

import (
	"fmt"
	"github.com/rabbitmq/amqp091-go"
)

type Pool struct {
	connection  *amqp091.Connection   //the connection amqp
	maxChannels int                   //the maximum quantity of channels of pool
	channels    chan *amqp091.Channel //a go channel to store the channels
	logger      logger                // a logger to log information
}

// NewPool create a new pool of channels
func NewPool(connectionString string, maxChannels int, logger logger) (*Pool, error) {

	connection, err := connect(connectionString)
	if err != nil {
		return nil, err
	}
	reusableChannels := make(chan *amqp091.Channel, maxChannels)

	for id := 0; id < maxChannels; id++ {
		reusableChannel, err := newChannel(connection)
		if err != nil {
			return nil, err
		}

		reusableChannels <- reusableChannel
	}

	logger.Info("Pool created successfully")
	return &Pool{
		connection:  connection,
		maxChannels: maxChannels,
		channels:    reusableChannels,
		logger:      logger,
	}, nil
}

// Close the connection with the broker amqp
func (pool *Pool) Close() error {

	connection := pool.connection

	if err := connection.Close(); err != nil {
		errMsg := fmt.Sprintf(
			"Occurred an error to try close the connection with the amqp broker: %v",
			err.Error(),
		)
		return fmt.Errorf(errMsg)
	}
	close(pool.channels)

	pool.logger.Info("Connection with the amqp broker was closed successfully")
	return nil
}

// GetChannel get a channel of the pool to use
func (pool *Pool) GetChannel() *ReusableChannel {
	return newReusableChannel(<-pool.channels, pool.channels)
}
