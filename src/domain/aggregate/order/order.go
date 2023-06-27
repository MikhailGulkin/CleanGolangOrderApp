package domain

import (
	"errors"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/consts"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/entities/address"
	domain "github.com/MikhailGulkin/simpleGoOrderApp/src/domain/entities/product"
	vo "github.com/MikhailGulkin/simpleGoOrderApp/src/domain/value_object"
	"github.com/google/uuid"
	"strconv"
	"time"
)

type priceOrder float64

type Order struct {
	vo.OrderId
	products        []domain.Product
	orderStatus     consts.OrderStatus
	paymentMethod   consts.PaymentMethod
	deliveryAddress address.Address
	totalPrice      priceOrder
	date            time.Time
	serialNumber    int
}

func (_ Order) create(products *[]domain.Product, deliveryAddress address.Address, previousSerialNumber int) (Order, error) {
	OrderId, idError := uuid.NewUUID()

	if idError != nil {
		return Order{}, errors.New("When OrderId was created error occurred: " + idError.Error())
	}
	serialNumber, serialError := getCurrentSerialNumber(previousSerialNumber)
	if serialError != nil {
		return Order{}, errors.New(serialError.Error())
	}

	return Order{
		OrderId:         vo.OrderId{Value: OrderId},
		products:        *products,
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
func (c *Order) getTotalPrice() priceOrder {
	var total priceOrder
	for _, product := range c.products {
		total += priceOrder(product.Price)
	}
	return total
}
