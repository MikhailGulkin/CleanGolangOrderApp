package domain

import (
	"errors"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/consts"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/entities/order"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/exceptions"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/vo"
	"strconv"
	"time"
)

type PriceOrder float64

type Order struct {
	vo.OrderID
	products        []order.OrderProduct
	client          order.OrderClient
	orderStatus     consts.OrderStatus
	paymentMethod   consts.PaymentMethod
	deliveryAddress order.OrderAddress
	totalPrice      PriceOrder
	date            time.Time
	serialNumber    int
	closed          bool
}

func (Order) Create(orderID vo.OrderID, deliveryAddress order.OrderAddress, client order.OrderClient, previousSerialNumber int) (Order, error) {
	serialNumber, serialError := getCurrentSerialNumber(previousSerialNumber)
	if serialError != nil {
		return Order{}, errors.New(serialError.Error())
	}

	return Order{
		OrderID:         orderID,
		orderStatus:     consts.New,
		client:          client,
		deliveryAddress: deliveryAddress,
		paymentMethod:   consts.Online,
		date:            time.Now(),
		serialNumber:    serialNumber,
	}, nil
}
func (o *Order) AddProduct(product order.OrderProduct) error {
	for _, p := range o.products {
		if p.ProductID == product.ProductID {
			exception := exceptions.ProductAlreadyContained{}.Exception(product.ProductID.String(), o.OrderID.Value.String())
			return &exception
		}
	}
	o.products = append(o.products, product)

	return nil
}
func (o *Order) RemoveProduct(product order.OrderProduct) error {
	start := -1
	for index, p := range o.products {
		if p.ProductID == product.ProductID {
			start = index
		}
	}
	if start == -1 {
		exception := exceptions.OrderProductNotExists{}.Exception(product.ProductID.String(), o.OrderID.Value.String())
		return &exception
	}
	o.products = append(o.products[:start], o.products[start+1:]...)

	return nil
}
func getCurrentSerialNumber(serialNumber int) (int, error) {
	if serialNumber > 100 || serialNumber < 1 {
		exception := exceptions.InvalidSerialNumber{}.Exception(strconv.Itoa(serialNumber))
		return -1, &exception
	}
	if serialNumber == 100 {
		return 1, nil
	}
	return serialNumber + 1, nil
}
func (o *Order) GetTotalPrice() PriceOrder {
	var total PriceOrder
	for _, orderProduct := range o.products {
		total += PriceOrder(orderProduct.Price)
	}
	return total
}
func (o *Order) GetSerialNumber() int {
	return o.serialNumber
}
func (o *Order) CheckNotClosed() error {
	if o.closed {
		exception := exceptions.OrderIsClosed{}.Exception(o.OrderID.Value.String())
		return &exception
	}
	return nil
}
func (o *Order) UpdateStatus(status consts.OrderStatus) error {
	if err := o.CheckNotClosed(); err != nil {
		return err
	}
	o.orderStatus = status

	if status == consts.Delivered || status == consts.Canceled {
		o.closed = true
	}
	return nil
}
