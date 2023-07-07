package convertors

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/dto"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/entities/product"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/value_object"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/models"
)

func ConvertProductModelToDTO(model models.Product) dto.Product {
	return dto.Product{
		ID:           model.ID,
		Name:         model.Name,
		Price:        model.Price,
		Discount:     model.Discount,
		Quantity:     model.Quantity,
		Description:  model.Description,
		Availability: model.Availability,
		CreatedAt:    model.CreatedAt,
	}
}
func ConvertProductModelsToDTOs(models []models.Product) []dto.Product {
	products := make([]dto.Product, len(models))
	for i, product := range models {
		products[i] = ConvertProductModelToDTO(product)
	}
	return products
}

func ConvertProductModelToEntity(model models.Product) product.Product {
	return product.Product{
		ProductID:    value_object.ProductID{Value: model.ID},
		Price:        model.Price,
		Name:         model.Name,
		Discount:     model.Discount,
		Quantity:     model.Quantity,
		Description:  model.Description,
		Availability: model.Availability,
	}
}
func ConvertProductEntityToModel(entity product.Product) models.Product {
	return models.Product{
		Base:         models.Base{ID: entity.ProductID.Value},
		Price:        entity.Price,
		Name:         entity.Name,
		Discount:     entity.Discount,
		Quantity:     entity.Quantity,
		Description:  entity.Description,
		Availability: entity.Availability,
	}
}
