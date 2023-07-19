package order

import (
	"errors"
	appRepo "github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/interfaces/persistence/repo"
	order "github.com/MikhailGulkin/simpleGoOrderApp/internal/domain/order/aggregate"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/db/models"
	repo "github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/db/repo"
	"gorm.io/gorm"
)

type RepoImpl struct {
	repo.BaseGormRepo
	appRepo.OrderRepo
}

func (repo *RepoImpl) AcquireLastOrder() (order.Order, error) {
	var orderModel models.Order
	result := repo.Session.
		Preload("Address").
		Preload("Client").
		Preload("Products").
		Order("created_at desc").
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

func (repo *RepoImpl) AddOrder(entity order.Order, tx any) error {
	model := ConvertOrderAggregateToModel(entity)
	result := tx.(*gorm.DB).
		Omit("Products.*").
		Create(&model)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
