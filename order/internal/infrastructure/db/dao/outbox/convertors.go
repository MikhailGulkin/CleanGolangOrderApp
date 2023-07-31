package outbox

import (
	"encoding/json"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/relay/dto"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/db/models"
)

func ConvertOutboxModelToDTO(model models.Outbox) (dto.Message, error) {
	b, err := json.Marshal(model.Payload)
	if err != nil {
		return dto.Message{}, err
	}
	return dto.Message{
		ID:       model.ID,
		Exchange: model.Exchange,
		Route:    model.Route,
		Payload:  b,
	}, nil
}
func ConvertOutboxModelsToDTOs(models []models.Outbox) []dto.Message {
	messages := make([]dto.Message, len(models))
	for _, model := range models {
		message, err := ConvertOutboxModelToDTO(model)
		if err != nil {
			continue
		}
		messages = append(messages, message)
	}
	return messages
}
