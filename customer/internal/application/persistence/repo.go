package persistence

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/common"
)

type Outbox interface {
	AddEvents(events []common.Event, tx interface{}) error
}
