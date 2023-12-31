package product

import (
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/domain/product/entities"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/domain/product/vo"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/db/models"
)

func ConvertProductModelToEntity(model models.Product) entities.Product {
	return entities.Product{
		ProductID:    vo.ProductID{Value: model.ID},
		Price:        vo.ProductPrice{Value: model.Price},
		Name:         model.Name,
		Discount:     vo.ProductDiscount{Value: model.Discount},
		Quantity:     model.Quantity,
		Description:  model.Description,
		Availability: model.Availability,
	}
}
func ConvertProductEntityToModel(entity entities.Product) models.Product {
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
