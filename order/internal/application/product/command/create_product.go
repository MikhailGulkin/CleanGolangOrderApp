package command

import (
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/common/interfaces/persistence"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/product/interfaces/command"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/product/interfaces/persistence/dao"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/domain/product/entities"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/domain/product/vo"
)

type CreateProductImpl struct {
	dao.ProductDAO
	persistence.UoW
	command.CreateProduct
}

func (interactor *CreateProductImpl) Create(command command.CreateProductCommand) error {
	price, priceErr := vo.ProductPrice{}.Create(command.Price)
	if priceErr != nil {
		return priceErr
	}
	discount, discountErr := vo.ProductDiscount{}.Create(command.Discount)
	if discountErr != nil {
		return discountErr
	}

	productEntity, err := entities.Product{}.Create(
		vo.ProductID{Value: command.ProductID},
		price,
		discount,
		command.Description,
		command.Name,
	)
	if err != nil {
		return err
	}
	err = interactor.ProductDAO.CreateProduct(productEntity, interactor.UoW.StartTx())
	if err != nil {
		return err
	}
	err = interactor.UoW.Commit()
	if err != nil {
		return err
	}

	return nil
}
