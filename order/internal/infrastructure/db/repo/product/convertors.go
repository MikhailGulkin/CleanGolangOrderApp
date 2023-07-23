package product

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/domain/product/aggregate"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/domain/product/vo"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/db/models"
)

func ConvertProductModelToEntity(model models.Product) aggregate.Product {
	return aggregate.Product{
		ProductID:    vo.ProductID{Value: model.ID},
		Price:        vo.ProductPrice{Value: model.Price},
		Name:         model.Name,
		Discount:     vo.ProductDiscount{Value: model.Discount},
		Quantity:     model.Quantity,
		Description:  model.Description,
		Availability: model.Availability,
	}
}
func ConvertProductEntityToModel(entity aggregate.Product) models.Product {
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
