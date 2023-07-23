package order

import (
	"fmt"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/common/consts"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/order/dto"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/order/interfaces/persistence/filters"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/order/interfaces/persistence/reader"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/db/models"
	r "github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/db/reader"
)

type ReaderImpl struct {
	reader.OrderReader
	r.BaseGormDAO
}

func (dao *ReaderImpl) GetAllOrders(filters filters.GetAllOrdersFilters) ([]dto.Order, error) {
	var orders []models.Order
	result := dao.Session.
		Preload("Products").
		Limit(int(filters.Limit)).
		Offset(int(filters.Offset)).
		Order(fmt.Sprintf("created_at %s", filters.Order)).
		Where("closed = ? AND saga_status = ?", false, consts.Approved).
		Find(&orders)
	if result.Error != nil {
		return []dto.Order{}, result.Error
	}
	return ConvertOrderModelsToDTOs(orders), nil
}
