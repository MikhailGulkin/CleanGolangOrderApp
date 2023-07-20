package query

import (
	base "github.com/MikhailGulkin/simpleGoOrderApp/internal/application/common/dto"
	baseFilters "github.com/MikhailGulkin/simpleGoOrderApp/internal/application/common/interfaces/persistence/filters"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/dto"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/interfaces/persistence/filters"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/interfaces/persistence/reader"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/interfaces/query"
)

type GetAllOrdersByUserIDImpl struct {
	query.GetAllOrdersByUserID
	reader.OrderCacheReader
}

func (interactor *GetAllOrdersByUserIDImpl) Get(q query.GetAllOrderByUserIDQuery) (dto.Orders, error) {
	orders, err := interactor.OrderCacheReader.GetAllOrdersByUserID(
		q.UserID,
		filters.GetAllOrdersByUserIDFilters{BaseFilters: baseFilters.BaseFilters{Limit: q.Limit, Offset: q.Offset, Order: q.Order}},
	)
	if err != nil {
		return dto.Orders{}, err
	}
	return dto.Orders{Orders: orders,
		BaseSequence: base.BaseSequence{Count: uint(len(orders)), Limit: q.Limit, Offset: q.Offset, Order: q.Order},
	}, err
}
