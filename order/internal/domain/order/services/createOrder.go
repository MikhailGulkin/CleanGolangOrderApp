package services

import (
	domain "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/domain/order/aggregate"
	order "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/domain/order/entities"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/domain/order/events"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/domain/order/exceptions"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/domain/order/vo"
	"github.com/google/uuid"
)

type Service struct {
}

func (Service) CreateOrder(
	orderID vo.OrderID,
	deliveryAddress uuid.UUID,
	client uuid.UUID,
	previousSerialNumber int,
	products []order.OrderProduct,
) (domain.Order, error) {
	createdOrder, orderError := domain.Order{}.Create(
		orderID,
		deliveryAddress,
		client,
		previousSerialNumber,
		products,
	)
	if orderError != nil {
		return domain.Order{}, orderError
	}
	if len(products) == 0 {
		orderException := exceptions.OrderProductsEmpty{}.Exception(orderID.ToString())
		return domain.Order{}, &orderException
	}
	productsEvent := make([]events.OrderCreateProductEvent, len(products))
	for index, product := range products {
		productsEvent[index] = events.OrderCreateProductEvent{
			ProductID:  product.ProductID,
			Name:       product.Name,
			TotalPrice: product.GetActualPrice(),
		}
	}
	createdOrder.RecordEvent(
		events.OrderCreated{}.Create(
			createdOrder.OrderID.Value,
			createdOrder.ClientID,
			string(createdOrder.PaymentMethod),
			createdOrder.SerialNumber,
			productsEvent,
			float64(createdOrder.GetTotalPrice()),
			createdOrder.DeliveryAddressID,
		),
	)
	createdOrder.RecordEvent(
		events.OrderCreateSaga{}.Create(
			createdOrder.OrderID.Value,
			float64(createdOrder.GetTotalPrice()),
			createdOrder.GetAllProductsIds(),
		),
	)
	return createdOrder, nil
}
