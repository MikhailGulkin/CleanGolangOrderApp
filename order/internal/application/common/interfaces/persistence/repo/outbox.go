package repo

import "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/domain/common/events"

type OutboxRepo interface {
	AddEvents(events []events.Event, tx interface{}) error
}
