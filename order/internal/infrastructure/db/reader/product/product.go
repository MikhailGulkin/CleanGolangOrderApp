package product

import (
	"errors"
	"fmt"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/product/dto"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/product/exceptions"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/product/interfaces/persistence/filters"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/product/interfaces/persistence/reader"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/db/models"
	db "github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/db/reader"
	"gorm.io/gorm"
)

type ReaderImpl struct {
	db.BaseGormDAO
	reader.ProductReader
}

func (dao *ReaderImpl) GetAllProducts(filters filters.GetAllProductsFilters) ([]dto.Product, error) {
	var products []models.Product
	result := dao.Session.
		Limit(int(filters.Limit)).
		Offset(int(filters.Offset)).
		Order(fmt.Sprintf("price %s", filters.Order)).
		Find(&products)
	if result.Error != nil {
		return []dto.Product{}, result.Error
	}
	return ConvertProductModelsToDTOs(products), nil
}
func (dao *ReaderImpl) GetProductByName(name string) (dto.Product, error) {
	var product models.Product
	result := dao.Session.Where("name = ?", name).First(&product)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		exception := exceptions.ProductNameNotExist{}.Exception(name)
		return dto.Product{}, &exception
	}
	if result.Error != nil {
		return dto.Product{}, result.Error
	}
	return ConvertProductModelToDTO(product), nil
}
