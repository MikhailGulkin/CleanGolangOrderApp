package query

import (
	base "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/common/dto"
	f "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/common/interfaces/persistence/filters"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/order/dto"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/order/interfaces/persistence/filters"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/order/interfaces/persistence/reader"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/order/interfaces/query"
)

type GetAllOrderImpl struct {
	query.GetAllOrders
	reader.OrderCacheReader
}

func (interactor *GetAllOrderImpl) Get(q query.GetAllOrderQuery) (dto.Orders, error) {
	orders, err := interactor.OrderCacheReader.GetAllOrders(
		filters.GetAllOrdersFilters{BaseFilters: f.BaseFilters{
			Limit: q.Limit, Offset: q.Offset, Order: q.Order,
		}},
	)
	if err != nil {
		return dto.Orders{}, err
	}
	return dto.Orders{Orders: orders, BaseSequence: base.BaseSequence{
		Count:  uint(len(orders)),
		Limit:  q.Limit,
		Offset: q.Offset,
		Order:  q.Order,
	}}, nil
}
