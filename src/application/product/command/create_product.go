package command

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/common/interfaces/persistence"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/product/interfaces/command"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/product/interfaces/persistence/repo"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/aggregate/product"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/vo"
)

type CreateProductImpl struct {
	repo.ProductRepo
	persistence.UoW
	command.CreateProduct
}

func (interactor *CreateProductImpl) Create(command command.CreateProductCommand) error {
	productEntity, err := product.Product{}.Create(
		vo.ProductID{Value: command.ProductID},
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
	err = interactor.UoW.Commit()
	if err != nil {
		return err
	}

	return nil
}
