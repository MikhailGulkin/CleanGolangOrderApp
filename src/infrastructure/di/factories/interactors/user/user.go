package user

import (
	r "github.com/MikhailGulkin/simpleGoOrderApp/src/application/address/interfaces/persistence/repo"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/common/interfaces/persistence"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/user/command"
	c "github.com/MikhailGulkin/simpleGoOrderApp/src/application/user/interfaces/command"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/user/interfaces/persistence/repo"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/services/user"
	"go.uber.org/fx"
)

func NewCreateUser(repo repo.UserRepo, uow persistence.UoW, repoAddress r.AddressRepo) c.CreateUser {
	return &command.CreateUserImpl{
		Service:     user.Service{},
		UoW:         uow,
		UserRepo:    repo,
		AddressRepo: repoAddress,
	}
}

var Module = fx.Provide(
	NewCreateUser,
)
