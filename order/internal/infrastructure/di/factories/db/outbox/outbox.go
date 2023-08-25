package outbox

import (
	appRepo "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/common/interfaces/persistence/repo"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/relay/interfaces/persistence/dao"
	relay "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/db/dao/outbox"
	repo "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/db/repo"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/db/repo/outbox"
	"go.uber.org/fx"
)

func BuildOutboxRepo(base repo.BaseGormRepo) appRepo.OutboxRepo {
	return &outbox.RepoImpl{BaseGormRepo: base}
}
func BuildOutboxDAO(base repo.BaseGormRepo) dao.OutboxDAO {
	return &relay.DAOImpl{BaseGormRepo: base}
}

var Module = fx.Provide(
	BuildOutboxRepo,
	BuildOutboxDAO,
)
