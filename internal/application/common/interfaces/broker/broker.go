package broker

import "context"

type MessageBroker interface {
	PublishMessage(ctx context.Context, exchangeName, routingKey string, message []byte) error
}
