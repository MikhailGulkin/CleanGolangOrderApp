package outbox

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/common/consts/outbox"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/relay/dto"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/relay/interfaces/persistence/dao"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/db/models"
	repo "github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/db/repo"
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
		Where("event_status = ?", outbox.Awaiting).
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
		UpdateColumn("event_status", outbox.Processed).Error
}
func (dao *DAOImpl) UpdateStatusMessagesByAggregateID(aggregateID uuid.UUID, status outbox.EventStatus, tx interface{}) error {
	return tx.(*gorm.DB).
		Model(&models.Outbox{}).
		Where("aggregate_id = ?", aggregateID).
		UpdateColumn("event_status", status).Error
}
