package order

import "github.com/google/uuid"

type OrderProduct struct {
	ProductID uuid.UUID
	Price     float64
}

func (OrderProduct) Create(productID uuid.UUID, price float64) (OrderProduct, error) {
	return OrderProduct{ProductID: productID, Price: price}, nil
}
