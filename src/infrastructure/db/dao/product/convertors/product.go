package convertors

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/dto"
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
