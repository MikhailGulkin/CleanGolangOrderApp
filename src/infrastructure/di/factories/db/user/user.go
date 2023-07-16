package user

import (
	appDAO "github.com/MikhailGulkin/simpleGoOrderApp/src/application/user/interfaces/persistence/dao"
	appRepo "github.com/MikhailGulkin/simpleGoOrderApp/src/application/user/interfaces/persistence/repo"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/dao/user"
	repo "github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/repo"
	userRepo "github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/repo/user"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

func BuildUserRepo(conn *gorm.DB) appRepo.UserRepo {
	return &userRepo.RepoImpl{BaseGormRepo: repo.BaseGormRepo{Session: conn}}
}
func BuildUserDAO(conn *gorm.DB) appDAO.UserDAO {
	return &user.DAOImpl{BaseGormRepo: repo.BaseGormRepo{Session: conn}}
}

var Module = fx.Provide(
	BuildUserRepo,
	BuildUserDAO,
)
