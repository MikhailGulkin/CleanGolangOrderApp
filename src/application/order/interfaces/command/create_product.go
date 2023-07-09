package command

import "github.com/google/uuid"

type CreateProductCommand struct {
	ProductID   uuid.UUID `json:"productID" binding:"required"`
	Price       float64   `json:"price" binding:"required"`
	Discount    int32     `json:"discount" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Name        string    `json:"name" binding:"required"`
}
type CreateProduct interface {
	Create(command CreateProductCommand) error
}
