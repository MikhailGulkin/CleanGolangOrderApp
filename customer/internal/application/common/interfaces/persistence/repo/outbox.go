package repo

import "github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/common/events"

type OutboxRepo interface {
	AddEvents(events []events.Event, tx interface{}) error
}
