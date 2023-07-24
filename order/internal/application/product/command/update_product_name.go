package command

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/common/interfaces/persistence"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/product/interfaces/command"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/product/interfaces/persistence/dao"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/domain/product/vo"
)

type UpdateProductNameImpl struct {
	dao.ProductDAO
	persistence.UoW
	command.UpdateProductName
}

func (interactor *UpdateProductNameImpl) Update(command command.UpdateProductNameCommand) error {
	productID := vo.ProductID{Value: command.ProductID}

	productEntity, err := interactor.ProductDAO.GetProductByID(productID)
	if err != nil {
		return err
	}
	err = productEntity.UpdateName(command.ProductName)
	if err != nil {
		return err
	}
	err = interactor.ProductDAO.UpdateProduct(productEntity, interactor.UoW.StartTx())
	if err != nil {
		return err
	}
	err = interactor.UoW.Commit()
	if err != nil {
		return err
	}
	return nil
}
