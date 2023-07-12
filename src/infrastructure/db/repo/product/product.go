package product

import (
	"errors"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/product/exceptions"
	appRepo "github.com/MikhailGulkin/simpleGoOrderApp/src/application/product/interfaces/persistence/repo"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/aggregate/product"
	product2 "github.com/MikhailGulkin/simpleGoOrderApp/src/domain/vo/product"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/models"
	repo "github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/repo"
	"gorm.io/gorm"
)

type RepoImpl struct {
	repo.BaseGormRepo
	appRepo.ProductRepo
}

func (repo *RepoImpl) AcquireProductByID(productID product2.ProductID) (product.Product, error) {
	var productModel models.Product
	result := repo.Session.Where("id = ?", productID.ToString()).First(&productModel)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		exception := exceptions.ProductIDNotExist{}.Exception(productID.ToString())
		return product.Product{}, &exception
	}
	if result.Error != nil {
		return product.Product{}, result.Error
	}
	return ConvertProductModelToEntity(productModel), nil
}
func (repo *RepoImpl) UpdateProduct(product product.Product, tx interface{}) error {
	productModel := ConvertProductEntityToModel(product)
	result := tx.(*gorm.DB).Where("id = ?", productModel.ID).Updates(productModel)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (repo *RepoImpl) AddProduct(product product.Product, tx interface{}) error {
	model := ConvertProductEntityToModel(product)
	return tx.(*gorm.DB).Create(&model).Error
}
