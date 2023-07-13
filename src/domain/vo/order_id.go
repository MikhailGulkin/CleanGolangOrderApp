package vo

import "github.com/google/uuid"

type OrderID struct {
	Value uuid.UUID
}

func (id OrderID) ToString() string {
	return id.Value.String()
}
