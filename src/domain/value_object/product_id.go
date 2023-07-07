package value_object

import "github.com/google/uuid"

type ProductID struct {
	Value uuid.UUID
}

func (id ProductID) ToString() string {
	return id.Value.String()
}
