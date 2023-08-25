package query

import (
	"fmt"
	base "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/common/dto"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/common/interfaces/logger"
	baseFilters "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/common/interfaces/persistence/filters"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/order/dto"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/order/interfaces/persistence/filters"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/order/interfaces/persistence/reader"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/order/interfaces/query"
)

type GetAllOrdersByUserIDImpl struct {
	query.GetAllOrdersByUserID
	reader.OrderCacheReader
	logger.Logger
}

func (interactor *GetAllOrdersByUserIDImpl) Get(q query.GetAllOrderByUserIDQuery) (dto.Orders, error) {
	orders, err := interactor.OrderCacheReader.GetAllOrdersByUserID(
		q.UserID,
		filters.GetAllOrdersByUserIDFilters{BaseFilters: baseFilters.BaseFilters{Limit: q.Limit, Offset: q.Offset, Order: q.Order}},
	)
	l := uint(len(orders))
	if err != nil {
		return dto.Orders{}, err
	}
	interactor.Logger.Info(
		fmt.Sprintf("Get all Orders by user id %s, count: %d, order: %s, limit: %d, offset: %d",
			q.UserID.String(), l, q.Order, q.Limit, q.Offset,
		))
	return dto.Orders{Orders: orders,
		BaseSequence: base.BaseSequence{Count: uint(len(orders)), Limit: q.Limit, Offset: q.Offset, Order: q.Order},
	}, err
}
