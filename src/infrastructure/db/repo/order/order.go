package order

import (
	"errors"
	appRepo "github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/persistence/repo"
	order "github.com/MikhailGulkin/simpleGoOrderApp/src/domain/aggregate/order"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/models"
	repo "github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/repo"
	"gorm.io/gorm"
)

type RepoImpl struct {
	repo.BaseGormRepo
	appRepo.OrderRepo
}

func (repo *RepoImpl) AcquireLastOrder() (order.Order, error) {
	var orderModel models.Order
	result := repo.Session.Preload("Address").Preload("Client").Preload("Products").Last(&orderModel)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return order.Order{}, nil
	}
	if result.Error != nil {
		return order.Order{}, result.Error
	}
	return ConvertOrderModelToAggregate(orderModel), nil
}

func (repo *RepoImpl) AddOrder(entity order.Order, tx any) error {
	products := make([]models.Product, len(entity.Products))
	for index, product := range entity.Products {
		products[index] = models.Product{
			Base: models.Base{ID: product.ProductID},
		}
	}
	model := ConvertOrderAggregateToModel(entity)
	result := tx.(*gorm.DB).Create(&model)
	if result.Error != nil {
		return result.Error
	}
	err := tx.(*gorm.DB).Model(&model).Association("Products").Append(products)
	if err != nil {
		return err
	}
	err = tx.(*gorm.DB).Model(&models.User{Base: models.Base{ID: entity.Client.ClientID}}).Association("Orders").Append(
		&model,
	)
	if err != nil {
		return err
	}
	return nil
}
