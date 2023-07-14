package order

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/dto"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/models"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/reader"
)

func ConvertOrderModelToDTO(model models.Order) dto.Order {
	products := make([]dto.Product, len(model.Products))
	for index, product := range model.Products {
		products[index] = dto.Product{
			ProductID:   product.ID,
			Name:        product.Name,
			ActualPrice: reader.CalculateTotalProductPrice(product),
		}
	}
	return dto.Order{
		OrderID:       model.ID,
		OrderStatus:   model.OrderStatus,
		PaymentMethod: model.PaymentMethod,
		Client: dto.Client{
			ClientID:   model.Client.ID,
			ClientName: model.Client.Username,
		},
		Address:      model.Address.GetFullAddress(),
		SerialNumber: model.SerialNumber,
		Products:     products,
		CreatedAt:    model.CreatedAt,
		TotalPrice:   reader.CalculateTotalOrderPrice(model.Products),
	}
}
func ConvertOrderModelsToDTOs(models []models.Order) []dto.Order {
	orders := make([]dto.Order, len(models))
	for index, model := range models {
		orders[index] = ConvertOrderModelToDTO(model)
	}
	return orders
}
