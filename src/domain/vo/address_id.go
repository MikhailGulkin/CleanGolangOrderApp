package vo

import "github.com/google/uuid"

type AddressID struct {
	Value uuid.UUID
}

func (id AddressID) ToString() string {
	return id.Value.String()
}
