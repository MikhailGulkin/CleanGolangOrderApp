package order

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/message_broker/broker_configurate/interfaces"
)

var ExchangeName = "Orders"
var QueueNameCreate = "OrdersCreate"
var QueueNameAddProduct = "OrdersAddProduct"

type BrokerSetup struct {
	interfaces.BaseMessageBrokerConfigure
}

func (b BrokerSetup) Setup() {
	b.DeclareExchange(ExchangeName)
	b.DeclareQueue(QueueNameCreate)
	b.DeclareQueue(QueueNameAddProduct)
	b.BindExchangeQueue(ExchangeName, "order_create", QueueNameCreate)
	b.BindExchangeQueue(ExchangeName, "order_add_product", QueueNameAddProduct)
}
func NewSetupBroker(broker interfaces.BaseMessageBrokerConfigure) BrokerSetup {
	return BrokerSetup{
		BaseMessageBrokerConfigure: broker,
	}
}