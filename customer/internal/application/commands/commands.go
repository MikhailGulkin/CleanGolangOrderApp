package commands

import (
	"github.com/google/uuid"
)

type CreateCustomerCommand struct {
	CustomerID uuid.UUID `json:"customerID" binding:"required"`
	FirstName  string    `json:"firstName" binding:"required"`
	LastName   string    `json:"lastName" binding:"required"`
	MiddleName string    `json:"middleName" binding:"required"`
	AddressID  uuid.UUID `json:"addressID" binding:"required"`
	Email      string    `json:"email" binding:"required"`
}
type UploadAvatarCommand struct {
	CustomerID uuid.UUID `json:"customerID" binding:"required"`
	Avatar     []byte    `json:"avatar" binding:"required"`
}
