package rabbit

import (
	"fmt"
	"github.com/rabbitmq/amqp091-go"
)

type Pool struct {
	connection                  *amqp091.Connection   //the connection amqp
	maxChannels                 int                   //the maximum quantity of channels of pool
	connectionCloseNotification chan *amqp091.Error   //a go channel to listen when the connection amqp was closed
	channels                    chan *amqp091.Channel //a go channel to store the channels
	channelsUsed                chan *amqp091.Channel //a go channel to store used channels
	logger                      Logger                // a logger to log information
}

// NewPool create a new pool of channels
func NewPool(connectionString string, maxChannels int, logger Logger) (*Pool, error) {
	connection, err := connect(connectionString)
	if err != nil {
		return nil, err
	}
	reusableChannels := make(chan *amqp091.Channel, maxChannels)

	for id := 1; id <= maxChannels; id++ {
		reusableChannel, err := newReusableChannel(connection)
		if err != nil {
			return nil, err
		}

		reusableChannels <- reusableChannel
	}

	connectionCloseNotification := make(chan *amqp091.Error)
	connectionCloseNotification = connection.NotifyClose(connectionCloseNotification)

	logger.Info("Pool created successfully")
	return &Pool{
		connection:                  connection,
		maxChannels:                 maxChannels,
		connectionCloseNotification: connectionCloseNotification,
		channels:                    reusableChannels,
		channelsUsed:                make(chan *amqp091.Channel, maxChannels),
		logger:                      logger,
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
	close(pool.channelsUsed)

	pool.logger.Info("Connection with the amqp broker was closed successfully")
	return nil
}

// GetChannel get a channel of the pool to use
func (pool *Pool) GetChannel() *amqp091.Channel {
	channel := <-pool.channels
	pool.channelsUsed <- channel
	return channel
}

// ReleaseChannel release a channel of the channelUsed pool
func (pool *Pool) ReleaseChannel() {
	channel := <-pool.channelsUsed
	pool.channels <- channel
}
