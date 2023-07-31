package aggregate

import (
	"errors"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/domain/common/aggregate"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/domain/order/consts"
	order "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/domain/order/entities"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/domain/order/events"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/domain/order/exceptions"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/domain/order/vo"
	"github.com/google/uuid"
	"strconv"
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
			product.GetActualPrice(),
			float64(o.GetTotalPrice()),
			product.Name,
			o.ClientID,
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
func (o *Order) DeleteOrder() error {
	err := o.CheckNotClosed()
	if err != nil {
		return err
	}
	o.RecordEvent(
		events.OrderDeleted{}.Create(o.OrderID.Value),
	)
	return nil
}
func (o *Order) GetAllProductsIds() []uuid.UUID {
	ids := make([]uuid.UUID, len(o.Products))
	for index, product := range o.Products {
		ids[index] = product.ProductID
	}
	return ids
}
