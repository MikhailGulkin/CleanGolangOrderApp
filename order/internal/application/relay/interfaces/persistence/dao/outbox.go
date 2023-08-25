package dao

import (
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/common/consts/outbox"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/relay/dto"
	"github.com/google/uuid"
)

type OutboxDAO interface {
	GetAllNonProcessedMessages() ([]dto.Message, error)
	UpdateMessage([]uuid.UUID) error
	UpdateStatusMessagesByAggregateID(aggregateID uuid.UUID, status outbox.EventStatus, tx interface{}) error
}
