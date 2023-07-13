package user

import (
	appRepo "github.com/MikhailGulkin/simpleGoOrderApp/src/application/user/interfaces/persistence/repo"
	repo "github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/repo"
	userRepo "github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/repo/user"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

func BuildUserRepo(conn *gorm.DB) appRepo.UserRepo {
	return &userRepo.RepoImpl{BaseGormRepo: repo.BaseGormRepo{Session: conn}}
}

var Module = fx.Provide(
	BuildUserRepo,
)
