package order

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/message_broker/broker_configurate/interfaces"
)

var ExchangeName = "Orders"
var QueueName = "Orders"

type BrokerSetup struct {
	interfaces.BaseMessageBrokerConfigure
}

func (b BrokerSetup) Setup() {
	b.DeclareExchange(ExchangeName)
	b.DeclareQueue(QueueName)
	b.BindExchangeQueue(ExchangeName, "order_create", QueueName)
	b.BindExchangeQueue(ExchangeName, "order_add_product", QueueName)
}
func NewSetupBroker(broker interfaces.BaseMessageBrokerConfigure) BrokerSetup {
	return BrokerSetup{
		BaseMessageBrokerConfigure: broker,
	}
}
