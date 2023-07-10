package query

import "github.com/MikhailGulkin/simpleGoOrderApp/src/application/common/interfaces/persistence/filters"

type BaseListQueryParams struct {
	Offset uint
	Limit  uint
	Order  filters.BaseOrder
}
