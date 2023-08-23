package vo

import "github.com/google/uuid"

type CustomerID struct {
	Value uuid.UUID
}

// StringID returns string representation of CustomerID
func (id *CustomerID) StringID() string {
	return id.Value.String()
}

// SetID sets value of CustomerID
func (id *CustomerID) SetID(value uuid.UUID) {
	id.Value = value
}

// GetID returns value of CustomerID
func (id *CustomerID) GetID() uuid.UUID {
	return id.Value
}
