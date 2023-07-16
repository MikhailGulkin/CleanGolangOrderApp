package order

import (
	order "github.com/MikhailGulkin/simpleGoOrderApp/src/domain/order/entities"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/models"
)

func ConvertProductsModelsToOrderEntity(models []models.Product) []order.OrderProduct {
	products := make([]order.OrderProduct, len(models))
	for index, model := range models {
		products[index] = order.OrderProduct{
			ProductID: model.ID,
			Price:     model.Price,
			Discount:  model.Discount,
		}
	}
	return products
}
func ConvertUserModelToOrderClient(model models.User) order.OrderClient {
	return order.OrderClient{
		ClientID: model.ID,
		Username: model.Username,
	}
}
func ConvertAddressModelToOrderAddress(model models.Address) order.OrderAddress {
	return order.OrderAddress{
		AddressID:   model.ID,
		FullAddress: model.GetFullAddress(),
	}
}