package query

import "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/common/interfaces/persistence/filters"

type BaseListQueryParams struct {
	Offset uint
	Limit  uint
	Order  filters.BaseOrder
}
