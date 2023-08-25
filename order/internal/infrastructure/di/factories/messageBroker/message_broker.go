package messagebrokerconfig

import (
	messagebroker "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/messageBroker"
	brokerconfigurate "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/messageBroker/brokerConfigurate"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/messageBroker/brokerConfigurate/order"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/messageBroker/controller"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	messagebroker.BuildAMPQ,
	brokerconfigurate.NewMessageBrokerConfigure,
	order.NewSetupBroker,
	brokerconfigurate.NewBrokerSetup,
	controller.NewMessageBroker,
)
