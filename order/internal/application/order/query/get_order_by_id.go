package query

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/dto"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/interfaces/persistence/reader"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/interfaces/query"
)

type GetOrderByIDImpl struct {
	reader.OrderReader
	query.GetOrderByID
}

func (interactor *GetOrderByIDImpl) Get(query query.GetOrderByIDQuery) (dto.Order, error) {
	order, err := interactor.OrderReader.GetOrderByID(query.ID)
	if err != nil {
		return dto.Order{}, err
	}
	return order, nil
}
