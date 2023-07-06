package reader

import (
	"errors"
	"fmt"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/dto"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/exceptions"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/persistence/dao"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/persistence/filters"
	db "github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/dao"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/dao/product/convertors"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/models"
	"gorm.io/gorm"
)

type ProductReaderImpl struct {
	db.BaseGormDAO
	dao.ProductReader
}

func (dao *ProductReaderImpl) GetAllProducts(filters filters.BaseFilters) ([]dto.Product, error) {
	var products []models.Product
	result := dao.Session.
		Limit(int(filters.Limit)).
		Offset(int(filters.Offset)).
		Order(fmt.Sprintf("price %s", filters.Order)).
		Find(&products)
	if result.Error != nil {
		return []dto.Product{}, result.Error
	}
	return convertors.ConvertProductModelsToDTOs(products), nil
}
func (dao *ProductReaderImpl) GetProductByName(name string) (dto.Product, error) {
	var product models.Product
	result := dao.Session.Where("name = ?", name).First(&product)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		exception := exceptions.ProductNameNotExist{}.Exception(name)
		return dto.Product{}, &exception
	}
	if result.Error != nil {
		return dto.Product{}, result.Error
	}
	fmt.Println(product)
	return convertors.ConvertProductModelToDTO(product), nil
}
