package interactors

import (
	"context"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/common/interfaces/broker"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/relay/interfaces/interactors"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/relay/interfaces/persistence/dao"
	"github.com/google/uuid"
	"time"
)

type RelayImpl struct {
	broker.MessageBroker
	dao.OutboxDAO
	interactors.Relay
}

func (r *RelayImpl) SendMessagesToBroker() {
	messages, err := r.GetAllNonProcessedMessages()
	if err != nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	processedIDs := make([]uuid.UUID, len(messages))
	for _, m := range messages {
		if err := r.PublishMessage(ctx, m.Exchange, m.Route, m.Payload); err != nil {
			continue
		}
		processedIDs = append(processedIDs, m.ID)
	}
	if len(processedIDs) != 0 {
		err = r.OutboxDAO.UpdateMessage(processedIDs)
		if err != nil {
			return
		}
	}
}
