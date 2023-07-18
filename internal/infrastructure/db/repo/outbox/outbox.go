package outbox

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/common/interfaces/persistence/repo"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/domain/common/events"
	base "github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/db/repo"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/db/repo/outbox/convertors"
	"gorm.io/gorm"
)

type RepoImpl struct {
	base.BaseGormRepo
	repo.OutboxRepo
}

func (repo *RepoImpl) AddEvents(events []events.Event, tx interface{}) error {
	converter, err := convertors.EventToOutbox{}.Create(events)
	if err != nil {
		return err
	}
	models := converter.Convert()
	result := tx.(*gorm.DB).Create(&models)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
