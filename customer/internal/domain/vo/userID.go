package vo

import "github.com/google/uuid"

type CustomerID struct {
	Value uuid.UUID
}

func (id *CustomerID) ToString() string {
	return id.Value.String()
}
func (id *CustomerID) SetID(value string) {
	id.Value = uuid.MustParse(value)
}
