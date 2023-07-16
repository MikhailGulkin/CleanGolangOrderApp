package order

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/persistence/dao"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/models"
	repo "github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/repo"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SagaDAOImpl struct {
	dao.OrderSagaDAO
	repo.BaseGormRepo
}

func (dao *SagaDAOImpl) UpdateOrderSagaStatus(orderID uuid.UUID, status string, tx interface{}) error {
	return tx.(*gorm.DB).
		Model(&models.Order{}).
		Where("id = ?", orderID).
		UpdateColumn("saga_status", status).Error
}
func (dao *SagaDAOImpl) DeleteOrderCascade(orderID uuid.UUID, tx interface{}) error {
	return tx.(*gorm.DB).
		Where("id = ?", orderID).
		Delete(&models.Order{}).Error
}
