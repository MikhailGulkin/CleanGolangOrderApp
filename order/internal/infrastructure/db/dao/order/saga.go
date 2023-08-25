package order

import (
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/common/consts"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/order/interfaces/persistence/dao"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/db/models"
	repo "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/db/repo"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SagaDAOImpl struct {
	dao.OrderSagaDAO
	repo.BaseGormRepo
}

func (dao *SagaDAOImpl) CheckSagaCompletion(orderID uuid.UUID) (bool, error) {
	var count int64
	res := dao.Session.Model(&models.Order{}).Where("saga_status = ? AND id = ?", consts.Pending, orderID).Count(&count)
	return count > 0, res.Error
}
func (dao *SagaDAOImpl) UpdateOrderSagaStatus(orderID uuid.UUID, status string, tx interface{}) error {
	return tx.(*gorm.DB).
		Model(&models.Order{}).
		Where("id = ?", orderID).
		UpdateColumn("saga_status", status).Error
}
func (dao *SagaDAOImpl) DeleteOrder(orderID uuid.UUID, tx interface{}) error {
	return tx.(*gorm.DB).
		Where("id = ?", orderID).
		Delete(&models.Order{}).Error
}
