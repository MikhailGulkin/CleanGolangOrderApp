package interfaces

type BaseMessageBrokerConfigure interface {
	DeclareExchange(exchangeName string)
	DeclareQueue(queueName string)
	BindExchangeQueue(exchangeName, key, queueName string)
}
