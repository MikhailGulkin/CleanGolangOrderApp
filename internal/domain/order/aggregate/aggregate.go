package aggregate

import (
	"errors"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/domain/common/aggregate"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/domain/order/consts"
	order "github.com/MikhailGulkin/simpleGoOrderApp/internal/domain/order/entities"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/domain/order/events"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/domain/order/exceptions"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/domain/order/vo"
	"strconv"
	"time"
)

type PriceOrder float64

type Order struct {
	aggregate.AggregateRoot
	vo.OrderID
	Products        []order.OrderProduct
	Client          order.OrderClient
	OrderStatus     consts.OrderStatus
	PaymentMethod   consts.PaymentMethod
	DeliveryAddress order.OrderAddress
	totalPrice      PriceOrder
	Date            time.Time
	SerialNumber    int
	Closed          bool
}

func (Order) Create(orderID vo.OrderID, deliveryAddress order.OrderAddress, client order.OrderClient, previousSerialNumber int) (Order, error) {
	serialNumber, serialError := getCurrentSerialNumber(previousSerialNumber)
	if serialError != nil {
		return Order{}, errors.New(serialError.Error())
	}
	return Order{
		OrderID:         orderID,
		OrderStatus:     consts.New,
		Client:          client,
		DeliveryAddress: deliveryAddress,
		PaymentMethod:   consts.Online,
		Date:            time.Now(),
		SerialNumber:    serialNumber,
	}, nil
}
func (o *Order) AddProduct(product order.OrderProduct) error {
	for _, p := range o.Products {
		if p.ProductID == product.ProductID {
			exception := exceptions.ProductAlreadyContained{}.Exception(product.ProductID.String(), o.OrderID.Value.String())
			return &exception
		}
	}
	o.Products = append(o.Products, product)
	o.RecordEvent(
		events.OrderAddProduct{}.Create(
			o.OrderID.Value,
			product.ProductID,
			float64(o.GetTotalPrice()),
			product.Price,
			o.Client.ClientID,
		),
	)
	return nil
}
func (o *Order) RemoveProduct(product order.OrderProduct) error {
	start := -1
	for index, p := range o.Products {
		if p.ProductID == product.ProductID {
			start = index
		}
	}
	if start == -1 {
		exception := exceptions.OrderProductNotExists{}.Exception(product.ProductID.String(), o.OrderID.Value.String())
		return &exception
	}
	o.Products = append(o.Products[:start], o.Products[start+1:]...)

	return nil
}
func getCurrentSerialNumber(serialNumber int) (int, error) {
	if serialNumber > 100 || serialNumber < 0 {
		exception := exceptions.InvalidSerialNumber{}.Exception(strconv.Itoa(serialNumber))
		return -1, &exception
	}
	if serialNumber == 100 {
		return 1, nil
	}
	return serialNumber + 1, nil
}
func (o *Order) GetTotalPrice() PriceOrder {
	var totalPrice float64
	var totalDiscount float64
	for _, orderProduct := range o.Products {
		totalPrice += orderProduct.Price
		totalDiscount += float64(orderProduct.Discount)
	}
	o.totalPrice = PriceOrder(totalPrice - ((totalDiscount / 100) * 100))
	return o.totalPrice
}
func (o *Order) GetSerialNumber() int {
	return o.SerialNumber
}
func (o *Order) CheckNotClosed() error {
	if o.Closed {
		exception := exceptions.OrderIsClosed{}.Exception(o.OrderID.ToString())
		return &exception
	}
	return nil
}
func (o *Order) UpdateStatus(status consts.OrderStatus) error {
	if err := o.CheckNotClosed(); err != nil {
		return err
	}
	o.OrderStatus = status

	if status == consts.Delivered || status == consts.Canceled {
		o.Closed = true
	}
	return nil
}
