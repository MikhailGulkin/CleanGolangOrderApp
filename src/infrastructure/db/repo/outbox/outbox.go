package outbox

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/common/interfaces/persistence/repo"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/common/events"
)

type OutboxRepoImpl struct {
	repo.OutboxRepo
}

func (repo *OutboxRepoImpl) AddEvents(events []events.Event, tx interface{}) error {
	return nil
}
