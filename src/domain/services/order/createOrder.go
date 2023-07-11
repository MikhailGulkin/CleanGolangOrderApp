package order

import (
	domain "github.com/MikhailGulkin/simpleGoOrderApp/src/domain/aggregate/order"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/aggregate/product"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/entities/order"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/vo"
)

func CreateOrder(orderID vo.OrderID, deliveryAddress order.OrderAddress, client order.OrderClient, previousSerialNumber int, products []product.Product) (domain.Order, error) {
	createdOrder, orderError := domain.Order{}.Create(
		orderID,
		deliveryAddress,
		client,
		previousSerialNumber,
	)
	if orderError != nil {
		return domain.Order{}, orderError
	}
	for _, p := range products {
		orderProduct, err := order.OrderProduct{}.Create(p.ProductID.Value, p.Price)
		if err != nil {
			return domain.Order{}, err
		}
		err = createdOrder.AddProduct(orderProduct)
		if err != nil {
			return domain.Order{}, err
		}
	}
	return createdOrder, nil
}