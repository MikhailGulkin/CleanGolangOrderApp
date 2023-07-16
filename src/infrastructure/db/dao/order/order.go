package order

import (
	"errors"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/exceptions"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/persistence/dao"
	order "github.com/MikhailGulkin/simpleGoOrderApp/src/domain/order/entities"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/models"
	repo "github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/repo"
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
func (dao *DAOImpl) GetClientByID(userID uuid.UUID) (order.OrderClient, error) {
	var userModel models.User
	result := dao.Session.Where("id = ?", userID.String()).First(&userModel)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		exception := exceptions.OrderClientIDNotExist{}.Exception(userID.String())
		return order.OrderClient{}, &exception
	}
	if result.Error != nil {
		return order.OrderClient{}, result.Error
	}
	return ConvertUserModelToOrderClient(userModel), nil
}
func (dao *DAOImpl) GetAddressByID(addressID uuid.UUID) (order.OrderAddress, error) {
	var addressModel models.Address
	result := dao.Session.Where("id = ?", addressID.String()).First(&addressModel)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		exception := exceptions.OrderAddressIDNotExist{}.Exception(addressID.String())
		return order.OrderAddress{}, &exception
	}
	if result.Error != nil {
		return order.OrderAddress{}, result.Error
	}
	return ConvertAddressModelToOrderAddress(addressModel), nil
}
