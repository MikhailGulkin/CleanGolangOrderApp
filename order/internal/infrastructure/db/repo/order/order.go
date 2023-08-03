package order

import (
	"errors"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/common/consts"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/order/exceptions"
	appRepo "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/order/interfaces/persistence/repo"
	order "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/domain/order/aggregate"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/db/models"
	repo "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/db/repo"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RepoImpl struct {
	repo.BaseGormRepo
	appRepo.OrderRepo
}

func (repo *RepoImpl) AcquireLastOrder() (order.Order, error) {
	var orderModel models.Order
	result := repo.Session.
		Preload("Products").
		Order("created_at desc").
		Where("saga_status = ?", consts.Approved).
		Where("deleted = ?", false).
		Limit(1).
		Find(&orderModel)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return order.Order{}, nil
	}
	if result.Error != nil {
		return order.Order{}, result.Error
	}
	return ConvertOrderModelToAggregate(orderModel), nil
}
func (repo *RepoImpl) AcquiredOrder(orderID uuid.UUID) (order.Order, error) {
	var orderModel models.Order
	result := repo.Session.
		Preload("Products").
		Where("saga_status = ? AND id = ?", consts.Approved, orderID).
		Where("deleted = ?", false).
		Find(&orderModel)
	err := exceptions.OrderIDNotExist{}.Exception(orderID.String())

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return order.Order{}, &err
	}
	if result.RowsAffected == 0 {
		return order.Order{}, &err
	}
	if result.Error != nil {
		return order.Order{}, result.Error
	}
	return ConvertOrderModelToAggregate(orderModel), nil
}

func (repo *RepoImpl) AddOrder(entity *order.Order, tx any) error {
	model := ConvertOrderAggregateToModel(entity)
	result := tx.(*gorm.DB).
		Omit("Products.*").
		Create(&model)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (repo *RepoImpl) UpdateOrder(entity *order.Order, tx any) error {
	model := ConvertOrderAggregateToModel(entity)
	return tx.(*gorm.DB).Updates(&model).Error
}
