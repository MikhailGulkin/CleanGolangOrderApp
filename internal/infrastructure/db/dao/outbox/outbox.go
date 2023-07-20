package outbox

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/relay/dto"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/relay/interfaces/persistence/dao"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/db/models"
	repo "github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/db/repo"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DAOImpl struct {
	dao.OutboxDAO
	repo.BaseGormRepo
}

func (dao *DAOImpl) GetAllNonProcessedMessages() ([]dto.Message, error) {
	var messages []models.Outbox
	result := dao.Session.
		Where("event_status = ?", models.Awaiting).
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
		UpdateColumn("event_status", models.Processed).Error
}
func (dao *DAOImpl) UpdateStatusMessagesByAggregateID(aggregateID uuid.UUID, status int, tx interface{}) error {
	return tx.(*gorm.DB).
		Model(&models.Outbox{}).
		Where("aggregate_id = ?", aggregateID).
		UpdateColumn("event_status", models.OutboxEventStatus(status)).Error
}
