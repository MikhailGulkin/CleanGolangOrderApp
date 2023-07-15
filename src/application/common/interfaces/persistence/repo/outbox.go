package repo

import "github.com/MikhailGulkin/simpleGoOrderApp/src/domain/common/events"

type OutBoxRepo interface {
	AddEvents(events []events.Event, tx interface{}) error
}
