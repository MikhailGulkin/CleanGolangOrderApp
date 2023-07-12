package product

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/aggregate/product"
	vo "github.com/MikhailGulkin/simpleGoOrderApp/src/domain/vo/product"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/models"
)

func ConvertProductModelToEntity(model models.Product) product.Product {
	return product.Product{
		ProductID:    vo.ProductID{Value: model.ID},
		Price:        vo.ProductPrice{Value: model.Price},
		Name:         model.Name,
		Discount:     vo.ProductDiscount{Value: model.Discount},
		Quantity:     model.Quantity,
		Description:  model.Description,
		Availability: model.Availability,
	}
}
func ConvertProductEntityToModel(entity product.Product) models.Product {
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
