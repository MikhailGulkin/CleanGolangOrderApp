package utils

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/aggregate/product"
	vo "github.com/MikhailGulkin/simpleGoOrderApp/src/domain/vo/product"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateValidProductEntity() product.Product {
	entity, _ := product.Product{}.Create(vo.ProductID{Value: uuid.New()}, vo.ProductPrice{Value: 100}, vo.ProductDiscount{Value: 10}, "", "some name")
	return entity
}
func CreateValidProductModel(entity product.Product) models.Product {
	return models.Product{
		Base:         models.Base{ID: entity.ProductID.Value},
		Price:        entity.Price.Value,
		Name:         entity.Name,
		Discount:     entity.Discount.Value,
		Quantity:     entity.Quantity,
		Description:  entity.Description,
		Availability: entity.Availability,
	}
}
func CreateProductInDB(model *models.Product, db *gorm.DB) {
	db.Create(&model)
}
