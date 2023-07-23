package order

import (
	order "github.com/MikhailGulkin/simpleGoOrderApp/order/internal/domain/order/entities"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/db/models"
)

func ConvertProductsModelsToOrderEntity(models []models.Product) []order.OrderProduct {
	products := make([]order.OrderProduct, len(models))
	for index, model := range models {
		products[index] = order.OrderProduct{
			ProductID: model.ID,
			Price:     model.Price,
			Name:      model.Name,
			Discount:  model.Discount,
		}
	}
	return products
}
