package aggregate

import (
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/domain/order/consts"
	order "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/domain/order/entities"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/domain/order/events"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/domain/order/exceptions"
	"time"
)

func (o *Order) AddProduct(product order.OrderProduct) error {
	if err := o.PreprocessOrder(); err != nil {
		return err
	}

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
	if err := o.PreprocessOrder(); err != nil {
		return err
	}

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

func (o *Order) UpdateStatus(status consts.OrderStatus) error {
	if err := o.PreprocessOrder(); err != nil {
		return err
	}
	o.OrderStatus = status

	if status == consts.Delivered || status == consts.Canceled {
		o.Closed = true
	}
	return nil
}
func (o *Order) DeleteOrder() error {
	if err := o.PreprocessOrder(); err != nil {
		return err
	}
	o.RecordEvent(
		events.OrderDeleted{}.Create(o.OrderID.Value),
	)
	t := time.Now()

	o.DeleteAt = &t
	o.Deleted = true

	return nil
}
