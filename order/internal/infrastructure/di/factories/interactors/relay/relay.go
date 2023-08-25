package relay

import (
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/common/interfaces/broker"
	impl "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/relay/interactors"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/relay/interfaces/interactors"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/relay/interfaces/persistence/dao"
	"go.uber.org/fx"
)

func NewOutboxRelay(dao dao.OutboxDAO, broker broker.MessageBroker) interactors.Relay {
	return &impl.RelayImpl{
		OutboxDAO:     dao,
		MessageBroker: broker,
	}
}

var Module = fx.Provide(
	NewOutboxRelay,
)
