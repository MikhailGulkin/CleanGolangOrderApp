package persistence

import "github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/common"

type EventStore interface {
	Save(customer common.Aggregate, tx interface{}) error
}
