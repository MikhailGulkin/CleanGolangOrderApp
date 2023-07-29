package query

import "github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/application/common/interfaces/persistence/filters"

type BaseListQueryParams struct {
	Offset uint
	Limit  uint
	Order  filters.BaseOrder
}
