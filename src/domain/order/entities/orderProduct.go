package order

import "github.com/google/uuid"

type OrderProduct struct {
	ProductID uuid.UUID
	Price     float64
	Discount  int32
}

func (OrderProduct) Create(productID uuid.UUID, price float64, discount int32) (OrderProduct, error) {
	return OrderProduct{ProductID: productID, Price: price, Discount: discount}, nil
}
