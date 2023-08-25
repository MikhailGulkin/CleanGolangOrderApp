package query

import (
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/product/dto"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/product/interfaces/persistence/reader"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/product/interfaces/query"
)

type GetProductByNameImpl struct {
	DAO reader.ProductReader
	query.GetProductByName
}

func (interactor *GetProductByNameImpl) Get(query query.GetProductByNameQuery) (dto.Product, error) {
	productByName, err := interactor.DAO.GetProductByName(query.Name)
	if err != nil {
		return dto.Product{}, err
	}
	return productByName, nil
}
