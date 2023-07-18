package address

import (
	appRepo "github.com/MikhailGulkin/simpleGoOrderApp/internal/application/address/interfaces/persistence/repo"
	repo "github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/db/repo"
	addressRepo "github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/db/repo/address"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

func BuildAddressRepo(conn *gorm.DB) appRepo.AddressRepo {
	return &addressRepo.RepoImpl{BaseGormRepo: repo.BaseGormRepo{Session: conn}}
}

var Module = fx.Provide(
	BuildAddressRepo,
)
