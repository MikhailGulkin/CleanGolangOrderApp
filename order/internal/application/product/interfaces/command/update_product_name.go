package command

import "github.com/google/uuid"

type UpdateProductNameCommand struct {
	ProductID   uuid.UUID
	ProductName string `json:"productName"`
}
type UpdateProductName interface {
	Update(command UpdateProductNameCommand) error
}
