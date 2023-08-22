package vo

import "github.com/google/uuid"

type CustomerID struct {
	Value uuid.UUID
}

func (id *CustomerID) StringID() string {
	return id.Value.String()
}
func (id *CustomerID) SetID(value uuid.UUID) {
	id.Value = value
}
func (id *CustomerID) GetID() uuid.UUID {
	return id.Value
}
