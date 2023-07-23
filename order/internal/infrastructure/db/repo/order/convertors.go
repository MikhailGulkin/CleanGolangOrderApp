package order

import (
	order "github.com/MikhailGulkin/simpleGoOrderApp/order/internal/domain/order/aggregate"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/domain/order/consts"
	o "github.com/MikhailGulkin/simpleGoOrderApp/order/internal/domain/order/entities"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/domain/order/vo"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/db/models"
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
		OrderID:  vo.OrderID{Value: model.ID},
		Products: products,
		Client: o.OrderClient{
			ClientID: model.Client.ID,
			Username: model.Client.Username,
		},
		OrderStatus:   consts.OrderStatus(model.OrderStatus),
		PaymentMethod: consts.PaymentMethod(model.PaymentMethod),
		DeliveryAddress: o.OrderAddress{
			AddressID:   model.Address.ID,
			FullAddress: model.Address.GetFullAddress(),
		},
		Date:         model.Base.CreatedAt,
		SerialNumber: model.SerialNumber,
		Closed:       model.Closed,
	}
}
func ConvertOrderAggregateToModel(order order.Order) models.Order {
	products := make([]models.Product, len(order.Products))
	for index, product := range order.Products {
		products[index] = models.Product{
			Base: models.Base{ID: product.ProductID},
		}
	}
	return models.Order{
		Base:          models.Base{ID: order.OrderID.Value, CreatedAt: order.Date},
		OrderStatus:   string(order.OrderStatus),
		ClientID:      order.Client.ClientID,
		PaymentMethod: string(order.PaymentMethod),
		AddressID:     order.DeliveryAddress.AddressID,
		Closed:        order.Closed,
		SerialNumber:  order.SerialNumber,
		Products:      products,
	}
}
