package order

import (
	"errors"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/order/exceptions"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/order/interfaces/persistence/dao"
	order "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/domain/order/entities"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/db/models"
	repo "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/db/repo"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DAOImpl struct {
	dao.OrderDAO
	repo.BaseGormRepo
}

func (dao *DAOImpl) GetProductsByIDs(productIDs []uuid.UUID) ([]order.OrderProduct, error) {
	ids := make([]string, len(productIDs))
	for index, productID := range productIDs {
		ids[index] = productID.String()
	}
	var productsModel []models.Product
	result := dao.Session.Where("id IN ?", productIDs).Find(&productsModel)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		exception := exceptions.ProductIDsNotExist{}.Exception(ids)
		return []order.OrderProduct{}, &exception
	}
	if result.Error != nil {
		return []order.OrderProduct{}, result.Error
	}
	return ConvertProductsModelsToOrderEntity(productsModel), nil
}

func (dao *DAOImpl) DeleteOrder(orderID uuid.UUID, tx interface{}) error {
	return tx.(*gorm.DB).
		Where("id = ?", orderID).
		Delete(&models.Order{}).Error
}
