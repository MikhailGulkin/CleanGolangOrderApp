package repo

import "github.com/MikhailGulkin/simpleGoOrderApp/internal/domain/common/events"

type OutboxRepo interface {
	AddEvents(events []events.Event, tx interface{}) error
}
