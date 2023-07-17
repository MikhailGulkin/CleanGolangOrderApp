package services

import (
	domain "github.com/MikhailGulkin/simpleGoOrderApp/src/domain/order/aggregate"
	order "github.com/MikhailGulkin/simpleGoOrderApp/src/domain/order/entities"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/order/events"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/order/exceptions"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/order/vo"
)

type Service struct {
}

func (Service) CreateOrder(
	orderID vo.OrderID,
	deliveryAddress order.OrderAddress,
	client order.OrderClient,
	previousSerialNumber int,
	products []order.OrderProduct,
) (domain.Order, error) {
	createdOrder, orderError := domain.Order{}.Create(
		orderID,
		deliveryAddress,
		client,
		previousSerialNumber,
	)
	if orderError != nil {
		return domain.Order{}, orderError
	}
	if len(products) == 0 {
		orderException := exceptions.OrderProductsEmpty{}.Exception(orderID.ToString())
		return domain.Order{}, &orderException
	}
	for _, product := range products {
		err := createdOrder.AddProduct(product)
		if err != nil {
			return domain.Order{}, err
		}
	}
	createdOrder.RecordEvent(
		events.OrderCreated{}.Create(
			createdOrder.OrderID.Value,
			events.OrderCreatedClient{ClientID: client.ClientID, Username: client.Username},
			string(createdOrder.PaymentMethod),
			createdOrder.DeliveryAddress.AddressID,
			createdOrder.SerialNumber,
			float64(createdOrder.GetTotalPrice()),
		),
	)
	return createdOrder, nil
}
