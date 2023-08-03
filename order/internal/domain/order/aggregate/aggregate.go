package aggregate

import (
	"errors"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/domain/common/aggregate"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/domain/order/consts"
	order "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/domain/order/entities"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/domain/order/vo"
	"github.com/google/uuid"
)

type PriceOrder float64

type Order struct {
	aggregate.AggregateRoot
	vo.OrderID
	Products          []order.OrderProduct
	ClientID          uuid.UUID
	OrderStatus       consts.OrderStatus
	PaymentMethod     consts.PaymentMethod
	DeliveryAddressID uuid.UUID
	totalPrice        PriceOrder
	vo.OrderInfo
	vo.OrderDeleted
}

func (Order) Create(
	orderID vo.OrderID,
	deliveryAddress uuid.UUID,
	client uuid.UUID,
	previousSerialNumber int,
	products []order.OrderProduct,
) (Order, error) {
	serialNumber, serialError := getCurrentSerialNumber(previousSerialNumber)
	if serialError != nil {
		return Order{}, errors.New(serialError.Error())
	}
	return Order{
		OrderID:           orderID,
		OrderStatus:       consts.New,
		Products:          products,
		ClientID:          client,
		DeliveryAddressID: deliveryAddress,
		PaymentMethod:     consts.Online,
		OrderInfo:         vo.OrderInfo{}.Create(serialNumber),
		OrderDeleted:      vo.OrderDeleted{}.Create(),
	}, nil
}
