package models

import "github.com/google/uuid"

type User struct {
	Base
	Username  string
	AddressID uuid.UUID
	Address   Address
}
