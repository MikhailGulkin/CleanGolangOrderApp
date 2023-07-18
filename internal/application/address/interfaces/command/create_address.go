package command

import "github.com/google/uuid"

type CreateAddressCommand struct {
	AddressID      uuid.UUID `json:"addressID" binding:"required"`
	BuildingNumber int       `json:"buildingNumber" binding:"required"`
	StreetName     string    `json:"streetName" binding:"required"`
	City           string    `json:"city" binding:"required"`
	Country        string    `json:"country" binding:"required"`
}
type CreateAddress interface {
	Create(command CreateAddressCommand) error
}
