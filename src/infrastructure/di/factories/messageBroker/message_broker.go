package messagebroker

import (
	messagebroker "github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/messageBroker"
	brokerconfigurate "github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/messageBroker/brokerConfigurate"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/messageBroker/brokerConfigurate/order"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/messageBroker/controller"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	messagebroker.BuildDial,
	messagebroker.BuildChannel,
	brokerconfigurate.NewMessageBrokerConfigure,
	order.NewSetupBroker,
	brokerconfigurate.NewBrokerSetup,
	controller.NewMessageBroker,
)
