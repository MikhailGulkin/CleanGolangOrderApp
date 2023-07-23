package order

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/messageBroker/brokerConfigurate/interfaces"
)

var ExchangeName = "Orders"
var QueueName = "Orders"

type BrokerSetup struct {
	interfaces.BaseMessageBrokerConfigure
}

func (b BrokerSetup) Setup() {
	b.DeclareExchange(ExchangeName)

	b.DeclareQueue(QueueName)
	b.BindExchangeQueue(ExchangeName, "Order.*", QueueName)

	// Optional queue build only for monolith structure
	b.DeclareQueue("CustomerSaga")
	b.BindExchangeQueue(ExchangeName, "Customer.Saga.*", "CustomerSaga")
	b.DeclareQueue("OrderSaga")
	b.BindExchangeQueue(ExchangeName, "Order.Saga.*", "")
}
func NewSetupBroker(broker interfaces.BaseMessageBrokerConfigure) BrokerSetup {
	return BrokerSetup{
		BaseMessageBrokerConfigure: broker,
	}
}
