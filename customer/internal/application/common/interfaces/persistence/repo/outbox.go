package repo

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/common"
)

type OutboxRepo interface {
	AddEvents(events []common.Event, tx interface{}) error
}
