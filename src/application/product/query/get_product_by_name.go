package query

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/product/dto"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/product/interfaces/persistence/reader"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/product/interfaces/query"
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
