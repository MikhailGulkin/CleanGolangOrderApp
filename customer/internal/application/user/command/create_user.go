package command

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/common/interfaces/persistence"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/user/interfaces/command"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/user/interfaces/persistence/dao"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/user/interfaces/persistence/repo"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/domain/user/services"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/domain/user/vo"
)

type CreateUserImpl struct {
	services.Service
	persistence.UoW
	repo.UserRepo
	dao.UserDAO
	command.CreateUser
}

func (interactor *CreateUserImpl) Create(command command.CreateUserCommand) error {
	address, err := interactor.UserDAO.GetUserAddress(command.AddressID)
	if err != nil {
		return err
	}
	entity := interactor.Service.CreateUser(vo.UserID{Value: command.UserID}, command.Username, address)
	err = interactor.UserRepo.AddUser(entity, interactor.UoW.StartTx())
	if err != nil {
		return err
	}
	err = interactor.UoW.Commit()
	if err != nil {
		return err
	}
	return nil
}
