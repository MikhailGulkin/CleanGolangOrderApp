package product

import (
	"errors"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/product/exceptions"
	appRepo "github.com/MikhailGulkin/simpleGoOrderApp/src/application/product/interfaces/persistence/repo"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/common/id"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/product/aggregate"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/models"
	repo "github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/repo"
	"gorm.io/gorm"
)

type RepoImpl struct {
	repo.BaseGormRepo
	appRepo.ProductRepo
}

func (repo *RepoImpl) AcquireProductsByIDs(productIDs []id.ID) ([]aggregate.Product, error) {
	ids := make([]string, len(productIDs))
	for index, productID := range productIDs {
		ids[index] = productID.ToString()
	}
	var productsModel []models.Product
	result := repo.Session.Where("id IN ?", ids).Find(&productsModel)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		exception := exceptions.ProductIDsNotExist{}.Exception(ids)
		return []aggregate.Product{}, &exception
	}
	if result.Error != nil {
		return []aggregate.Product{}, result.Error
	}
	return ConvertProductsModelsToAggregate(&productsModel), nil
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
