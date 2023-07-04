package dao

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/persistence"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/entities/product"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/models"
	"gorm.io/gorm"
)

type ProductDAOImpl struct {
	BaseGormDAO
	persistence.ProductDAO
}

func (dao *ProductDAOImpl) Create(product product.Product) error {
	productModel := models.Product{
		Base:         models.Base{ID: product.ProductID.Value},
		Price:        product.Price,
		Name:         product.Name,
		Discount:     product.Discount,
		Quantity:     product.Quantity,
		Description:  product.Description,
		Availability: product.Availability,
	}
	result := dao.session.Create(&productModel)
	return result.Error
}

func BuildProductDAO(conn *gorm.DB) persistence.ProductDAO {
	return &ProductDAOImpl{BaseGormDAO: BaseGormDAO{conn}}
}
