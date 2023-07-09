package product

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/common/interfaces/persistence"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/command"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/persistence/dao"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/entities/product"
)

type CreateProductImpl struct {
	dao.ProductRepo
	persistence.UoW
	command.CreateProduct
}

func (interactor *CreateProductImpl) Create(command command.CreateProductCommand) error {
	productEntity, err := product.Product{}.Create(
		command.ProductID,
		command.Price,
		command.Discount,
		command.Description,
		command.Name,
	)
	if err != nil {
		return err
	}
	interactor.UoW.StartTx()
	err = interactor.ProductRepo.Create(productEntity, interactor.UoW.GetTx())
	if err != nil {
		return err
	}
	interactor.UoW.Commit()

	return nil
}
