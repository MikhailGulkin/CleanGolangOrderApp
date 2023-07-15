package messagebroker

import (
	messagebroker "github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/message_broker"
	brokerconfigurate "github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/message_broker/broker_configurate"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/message_broker/broker_configurate/order"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/message_broker/controller"
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
