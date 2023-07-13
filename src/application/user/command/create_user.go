package command

import (
	addressRepo "github.com/MikhailGulkin/simpleGoOrderApp/src/application/address/interfaces/persistence/repo"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/common/interfaces/persistence"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/user/interfaces/command"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/user/interfaces/persistence/repo"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/services/user"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/vo"
)

type CreateUserImpl struct {
	user.Service
	persistence.UoW
	repo.UserRepo
	addressRepo.AddressRepo
	command.CreateUser
}

func (interactor *CreateUserImpl) Create(command command.CreateUserCommand) error {
	address, err := interactor.AddressRepo.AcquireAddressByID(vo.AddressID{Value: command.AddressID})
	if err != nil {
		return err
	}
	interactor.StartTx()
	entity := interactor.Service.CreateUser(vo.UserID{Value: command.UserID}, command.Username, address)
	err = interactor.UserRepo.AddUser(entity, interactor.GetTx())
	if err != nil {
		return err
	}
	err = interactor.UoW.Commit()
	if err != nil {
		return err
	}
	return nil
}
