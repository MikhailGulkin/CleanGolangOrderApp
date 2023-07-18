package relay

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/common/interfaces/broker"
	impl "github.com/MikhailGulkin/simpleGoOrderApp/internal/application/relay/interactors"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/relay/interfaces/interactors"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/relay/interfaces/persistence/dao"
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
