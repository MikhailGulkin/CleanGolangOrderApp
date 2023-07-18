package brokerconfigurate

import "github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/messageBroker/brokerConfigurate/order"

type Brokers []Broker
type Broker interface {
	Setup()
}

func NewBrokerSetup(
	orderSetup order.BrokerSetup,
) Brokers {
	return Brokers{
		orderSetup,
	}
}
func (r Brokers) Setup() {
	for _, broker := range r {
		broker.Setup()
	}
}
