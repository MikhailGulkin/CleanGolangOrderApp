package order

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/messageBroker/brokerConfigurate/interfaces"
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
	b.BindExchangeQueue(ExchangeName, "Order.Create", QueueNameCreate)
	b.BindExchangeQueue(ExchangeName, "Order.AddProduct", QueueNameAddProduct)
}
func NewSetupBroker(broker interfaces.BaseMessageBrokerConfigure) BrokerSetup {
	return BrokerSetup{
		BaseMessageBrokerConfigure: broker,
	}
}
