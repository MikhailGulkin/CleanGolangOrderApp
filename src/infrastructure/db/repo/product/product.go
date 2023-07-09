package product

import (
	"errors"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/exceptions"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/persistence/dao"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/entities/product"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/value_object"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/models"
	repo "github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/repo"
	"gorm.io/gorm"
)

type RepoImpl struct {
	repo.BaseGormRepo
	dao.ProductRepo
}

func (dao *RepoImpl) AcquireProductByID(productID value_object.ProductID) (product.Product, error) {
	var productModel models.Product
	result := dao.Session.Where("id = ?", productID.ToString()).First(&productModel)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		exception := exceptions.ProductIDNotExist{}.Exception(productID.ToString())
		return product.Product{}, &exception
	}
	if result.Error != nil {
		return product.Product{}, result.Error
	}
	return ConvertProductModelToEntity(productModel), nil
}
func (dao *RepoImpl) UpdateProduct(product product.Product, tx interface{}) error {
	productModel := ConvertProductEntityToModel(product)
	result := tx.(*gorm.DB).Where("id = ?", productModel.ID).Updates(productModel)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (dao *RepoImpl) Create(product product.Product, tx interface{}) error {
	productModel := models.Product{
		Base:         models.Base{ID: product.ProductID.Value},
		Price:        product.Price,
		Name:         product.Name,
		Discount:     product.Discount,
		Quantity:     product.Quantity,
		Description:  product.Description,
		Availability: product.Availability,
	}
	return tx.(*gorm.DB).Create(&productModel).Error
}
