package vo

import "github.com/google/uuid"

type UserID struct {
	Value uuid.UUID
}

func (id *UserID) ToString() string {
	return id.Value.String()
}
func (id *UserID) SetID(value string) {
	id.Value = uuid.MustParse(value)
}
