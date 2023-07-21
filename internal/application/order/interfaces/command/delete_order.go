package command

import "github.com/google/uuid"

type DeleteOrderCommand struct {
	OrderID uuid.UUID `json:"orderID,omitempty"`
}
type DeleteOrder interface {
	Delete(command DeleteOrderCommand) error
}
