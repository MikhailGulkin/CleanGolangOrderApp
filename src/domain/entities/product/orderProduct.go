package product

import "github.com/google/uuid"

type OrderProduct struct {
	ProductID uuid.UUID
	Price     float64
}
