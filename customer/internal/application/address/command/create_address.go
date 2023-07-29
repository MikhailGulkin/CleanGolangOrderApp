package command

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/application/address/interfaces/command"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/application/address/interfaces/persistence/repo"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/application/common/interfaces/persistence"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/address/aggregate"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/address/vo"
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
	err := interactor.AddressRepo.AddAddress(addressEntity, interactor.UoW.StartTx())
	if err != nil {
		return err
	}
	err = interactor.UoW.Commit()
	if err != nil {
		return err
	}
	return nil
}
