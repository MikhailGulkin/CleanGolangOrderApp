package order

import "github.com/google/uuid"

type OrderProduct struct {
	ProductID uuid.UUID
	Name      string
	Price     float64
	Discount  int32
}

func (OrderProduct) Create(productID uuid.UUID, price float64, discount int32, name string) (OrderProduct, error) {
	return OrderProduct{ProductID: productID, Price: price, Discount: discount, Name: name}, nil
}
func (o *OrderProduct) GetActualPrice() float64 {
	return o.Price - ((float64(o.Discount) / 100) * 100)
}
