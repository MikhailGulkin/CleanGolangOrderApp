package rabbit

import "github.com/rabbitmq/amqp091-go"

type ReusableChannel struct {
	*amqp091.Channel                       // the Channel amqp
	channels         chan *amqp091.Channel // a go channel to store the channels
}

// newChannel create a new channel
func newChannel(connection *amqp091.Connection) (*amqp091.Channel, error) {
	channel, err := connection.Channel()
	if err != nil {
		return nil, err
	}
	return channel, nil
}

// newReusableChannel create a new reusable channel
func newReusableChannel(channel *amqp091.Channel, channels chan *amqp091.Channel) *ReusableChannel {
	return &ReusableChannel{channels: channels, Channel: channel}
}
func (r *ReusableChannel) Release() {
	r.channels <- r.Channel
}
