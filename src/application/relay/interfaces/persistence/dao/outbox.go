package dao

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/relay/dto"
	"github.com/google/uuid"
)

type OutboxDAO interface {
	GetAllNonProcessedMessages() ([]dto.Message, error)
	UpdateMessage([]uuid.UUID) error
}
