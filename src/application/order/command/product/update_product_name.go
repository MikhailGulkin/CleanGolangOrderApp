package product

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/common/interfaces/persistence"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/command"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/persistence/repo"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/value_object"
)

type UpdateProductNameImpl struct {
	repo.ProductRepo
	persistence.UoW
	command.UpdateProductName
}

func (interactor *UpdateProductNameImpl) Update(command command.UpdateProductNameCommand) error {
	productID := value_object.ProductID{Value: command.ProductID}

	productEntity, err := interactor.ProductRepo.AcquireProductByID(productID)
	if err != nil {
		return err
	}
	err = productEntity.UpdateName(command.ProductName)
	if err != nil {
		return err
	}
	interactor.UoW.StartTx()
	err = interactor.ProductRepo.UpdateProduct(productEntity, interactor.UoW.GetTx())
	if err != nil {
		return err
	}
	interactor.UoW.Commit()
	return nil
}
