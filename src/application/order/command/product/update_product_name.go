package product

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/common/interfaces/persistence"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/command"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/persistence/dao"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/value_object"
)

type UpdateProductNameImpl struct {
	dao.ProductDAO
	persistence.UoW
	command.UpdateProductName
}

func (interactor *UpdateProductNameImpl) Update(command command.UpdateProductNameCommand) error {
	productID := value_object.ProductID{Value: command.ProductID}

	productEntity, err := interactor.ProductDAO.AcquireProductByID(productID)
	if err != nil {
		return err
	}
	err = productEntity.UpdateName(command.ProductName)
	if err != nil {
		return err
	}
	interactor.StartTx()
	err = interactor.ProductDAO.UpdateProduct(productEntity, interactor.GetTx())
	if err != nil {
		return err
	}
	interactor.Commit()
	return nil
}
