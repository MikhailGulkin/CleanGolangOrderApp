package messagebroker

import (
	messagebroker "github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/messageBroker"
	brokerconfigurate "github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/messageBroker/brokerConfigurate"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/messageBroker/brokerConfigurate/order"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/messageBroker/controller"
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
