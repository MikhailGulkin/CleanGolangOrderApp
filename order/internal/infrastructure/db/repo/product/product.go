package product

import (
	"errors"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/product/exceptions"
	appRepo "github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/product/interfaces/persistence/repo"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/domain/common/id"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/domain/product/aggregate"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/db/models"
	repo "github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/db/repo"
	"gorm.io/gorm"
)

type RepoImpl struct {
	repo.BaseGormRepo
	appRepo.ProductRepo
}

func (repo *RepoImpl) AcquireProductByID(productID id.ID) (aggregate.Product, error) {
	var productModel models.Product
	result := repo.Session.Where("id = ?", productID.ToString()).First(&productModel)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		exception := exceptions.ProductIDNotExist{}.Exception(productID.ToString())
		return aggregate.Product{}, &exception
	}
	if result.Error != nil {
		return aggregate.Product{}, result.Error
	}
	return ConvertProductModelToEntity(productModel), nil
}
func (repo *RepoImpl) UpdateProduct(product aggregate.Product, tx interface{}) error {
	productModel := ConvertProductEntityToModel(product)
	result := tx.(*gorm.DB).Where("id = ?", productModel.ID).Updates(productModel)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (repo *RepoImpl) AddProduct(product aggregate.Product, tx interface{}) error {
	model := ConvertProductEntityToModel(product)
	return tx.(*gorm.DB).Create(&model).Error
}
