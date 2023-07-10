package domain

import (
	"errors"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/consts"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/entities/address"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/entities/product"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/vo"
	"github.com/google/uuid"
	"strconv"
	"time"
)

type PriceOrder float64

type Order struct {
	vo.OrderId
	products        []product.OrderProduct
	orderStatus     consts.OrderStatus
	paymentMethod   consts.PaymentMethod
	deliveryAddress address.Address
	totalPrice      PriceOrder
	date            time.Time
	serialNumber    int
}

func (Order) Create(orderID uuid.UUID, products []product.OrderProduct, deliveryAddress address.Address, previousSerialNumber int) (Order, error) {
	serialNumber, serialError := getCurrentSerialNumber(previousSerialNumber)
	if serialError != nil {
		return Order{}, errors.New(serialError.Error())
	}

	return Order{
		OrderId:         vo.OrderId{Value: orderID},
		products:        products,
		orderStatus:     consts.New,
		deliveryAddress: deliveryAddress,
		paymentMethod:   consts.Online,
		date:            time.Now(),
		serialNumber:    serialNumber,
	}, nil
}
func getCurrentSerialNumber(serialNumber int) (int, error) {
	if serialNumber > 100 || serialNumber < 1 {
		return -1, errors.New("wrong SerialNumber creation, serial: " + strconv.Itoa(serialNumber))
	}
	if serialNumber == 100 {
		return 1, nil
	}
	return serialNumber + 1, nil
}
func (c *Order) GetTotalPrice() PriceOrder {
	var total PriceOrder
	for _, orderProduct := range c.products {
		total += PriceOrder(orderProduct.Price)
	}
	return total
}
