package query

import (
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/order/dto"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/order/interfaces/persistence/reader"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/order/interfaces/query"
)

type GetOrderByIDImpl struct {
	reader.OrderCacheReader
	query.GetOrderByID
}

func (interactor *GetOrderByIDImpl) Get(query query.GetOrderByIDQuery) (dto.Order, error) {
	order, err := interactor.OrderCacheReader.GetOrderByID(query.ID)
	if err != nil {
		return dto.Order{}, err
	}
	return order, nil
}
