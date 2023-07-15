package relay

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/common/interfaces/broker"
	impl "github.com/MikhailGulkin/simpleGoOrderApp/src/application/relay/interactors"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/relay/interfaces/interactors"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/relay/interfaces/persistence/dao"
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
