package product

import (
	"errors"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/product/exceptions"
	appDAO "github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/product/interfaces/persistence/dao"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/domain/common/id"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/domain/product/entities"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/db/models"
	repo "github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/db/repo"
	"gorm.io/gorm"
)

type DAOImpl struct {
	repo.BaseGormRepo
	appDAO.ProductDAO
}

func (repo *DAOImpl) GetProductByID(productID id.ID) (entities.Product, error) {
	var productModel models.Product
	result := repo.Session.Where("id = ?", productID.ToString()).First(&productModel)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		exception := exceptions.ProductIDNotExist{}.Exception(productID.ToString())
		return entities.Product{}, &exception
	}
	if result.Error != nil {
		return entities.Product{}, result.Error
	}
	return ConvertProductModelToEntity(productModel), nil
}
func (repo *DAOImpl) UpdateProduct(product entities.Product, tx interface{}) error {
	productModel := ConvertProductEntityToModel(product)
	result := tx.(*gorm.DB).Where("id = ?", productModel.ID).Updates(productModel)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (repo *DAOImpl) CreateProduct(product entities.Product, tx interface{}) error {
	model := ConvertProductEntityToModel(product)
	return tx.(*gorm.DB).Create(&model).Error
}
