package outbox

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/relay/dto"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/relay/interfaces/persistence/dao"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/models"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/reader"
	"github.com/google/uuid"
)

type DAOImpl struct {
	dao.OutboxDAO
	reader.BaseGormDAO
}

func (dao *DAOImpl) GetAllNonProcessedMessages() ([]dto.Message, error) {
	var messages []models.Outbox
	result := dao.Session.
		Where("is_processed = ?", false).
		Find(&messages)
	if result.Error != nil {
		return []dto.Message{}, nil
	}
	return ConvertOutboxModelsToDTOs(messages), nil
}
func (dao *DAOImpl) UpdateMessage(ids []uuid.UUID) error {
	return dao.Session.
		Model(&models.Outbox{}).
		Where("id IN ?", ids).
		UpdateColumn("is_processed", true).Error
}