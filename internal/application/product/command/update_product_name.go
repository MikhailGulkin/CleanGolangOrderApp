package command

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/common/interfaces/persistence"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/product/interfaces/command"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/product/interfaces/persistence/repo"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/domain/product/vo"
)

type UpdateProductNameImpl struct {
	repo.ProductRepo
	persistence.UoW
	command.UpdateProductName
}

func (interactor *UpdateProductNameImpl) Update(command command.UpdateProductNameCommand) error {
	productID := vo.ProductID{Value: command.ProductID}

	productEntity, err := interactor.ProductRepo.AcquireProductByID(productID)
	if err != nil {
		return err
	}
	err = productEntity.UpdateName(command.ProductName)
	if err != nil {
		return err
	}
	err = interactor.ProductRepo.UpdateProduct(productEntity, interactor.UoW.StartTx())
	if err != nil {
		return err
	}
	err = interactor.UoW.Commit()
	if err != nil {
		return err
	}
	return nil
}