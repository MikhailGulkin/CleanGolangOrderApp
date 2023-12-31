package order

import (
	order "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/domain/order/aggregate"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/domain/order/consts"
	o "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/domain/order/entities"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/domain/order/vo"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/db/models"
	"time"
)

func ConvertOrderModelToAggregate(model models.Order) order.Order {
	products := make([]o.OrderProduct, len(model.Products))
	for index, product := range model.Products {
		products[index] = o.OrderProduct{
			ProductID: product.ID,
			Price:     product.Price,
			Discount:  product.Discount,
		}
	}
	return order.Order{
		OrderID:           vo.OrderID{Value: model.ID},
		Products:          products,
		ClientID:          model.ClientID,
		OrderStatus:       consts.OrderStatus(model.OrderStatus),
		PaymentMethod:     consts.PaymentMethod(model.PaymentMethod),
		DeliveryAddressID: model.AddressID,
		OrderInfo: vo.OrderInfo{
			Date:         model.Base.CreatedAt,
			SerialNumber: model.SerialNumber,
			Closed:       model.Closed,
		},
		OrderDeleted: vo.OrderDeleted{
			Deleted:  model.Deleted,
			DeleteAt: model.DeletedAt,
		},
	}
}
func ConvertOrderAggregateToModel(order *order.Order) models.Order {
	products := make([]models.Product, len(order.Products))
	for index, product := range order.Products {
		products[index] = models.Product{
			Base: models.Base{ID: product.ProductID},
		}
	}
	return models.Order{
		Base: models.Base{
			ID:        order.OrderID.Value,
			UpdatedAt: time.Now(),
			CreatedAt: order.Date,
			DeletedAt: order.DeleteAt,
		},
		OrderStatus:   string(order.OrderStatus),
		ClientID:      order.ClientID,
		PaymentMethod: string(order.PaymentMethod),
		AddressID:     order.DeliveryAddressID,
		Closed:        order.Closed,
		SerialNumber:  order.SerialNumber,
		Products:      products,
		Deleted:       order.Deleted,
	}
}
