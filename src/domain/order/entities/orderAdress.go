package order

import "github.com/google/uuid"

type OrderAddress struct {
	AddressID   uuid.UUID
	FullAddress string
}

func (OrderAddress) Create(id uuid.UUID, address string) (OrderAddress, error) {
	return OrderAddress{
		AddressID:   id,
		FullAddress: address,
	}, nil
}
