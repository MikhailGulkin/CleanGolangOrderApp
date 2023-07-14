package command

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/address/interfaces/command"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/address/interfaces/persistence/repo"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/common/interfaces/persistence"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/address/aggregate"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/address/vo"
)

type CreateAddressImpl struct {
	command.CreateAddress
	persistence.UoW
	repo.AddressRepo
}

func (interactor *CreateAddressImpl) Create(command command.CreateAddressCommand) error {
	addressEntity := aggregate.Address{}.Create(
		vo.AddressID{Value: command.AddressID},
		command.BuildingNumber,
		command.StreetName,
		command.City,
		command.Country,
	)
	interactor.UoW.StartTx()
	err := interactor.AddressRepo.AddAddress(addressEntity, interactor.UoW.GetTx())
	if err != nil {
		return err
	}
	err = interactor.UoW.Commit()
	if err != nil {
		return err
	}
	return nil
}
