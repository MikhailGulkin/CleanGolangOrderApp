package persistence

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/common"
)

type EventStore interface {
	Create(customer common.Aggregate, tx interface{}) error
	Update(customer common.Aggregate, tx interface{}) error
	Find(id string) (common.Aggregate, error)
	Exists(id string) error
}
