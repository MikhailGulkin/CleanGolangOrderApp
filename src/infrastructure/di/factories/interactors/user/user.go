package user

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/common/interfaces/persistence"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/user/command"
	c "github.com/MikhailGulkin/simpleGoOrderApp/src/application/user/interfaces/command"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/user/interfaces/persistence/dao"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/user/interfaces/persistence/repo"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/user/services"
	"go.uber.org/fx"
)

func NewCreateUser(repo repo.UserRepo, uow persistence.UoW, userDAO dao.UserDAO) c.CreateUser {
	return &command.CreateUserImpl{
		Service:  services.Service{},
		UoW:      uow,
		UserRepo: repo,
		UserDAO:  userDAO,
	}
}

var Module = fx.Provide(
	NewCreateUser,
)
