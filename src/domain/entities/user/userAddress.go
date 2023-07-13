package user

import "github.com/google/uuid"

type UserAddress struct {
	AddressID uuid.UUID
	Address   string
}

func (UserAddress) Create(id uuid.UUID, address string) UserAddress {
	return UserAddress{
		AddressID: id,
		Address:   address,
	}
}
