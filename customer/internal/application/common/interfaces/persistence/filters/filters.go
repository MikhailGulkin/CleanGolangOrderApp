package filters

type BaseOrder string

const (
	ASC  BaseOrder = "asc"
	DESC BaseOrder = "desc"
)

type BaseFilters struct {
	Limit  uint
	Offset uint
	Order  BaseOrder
}

func (filter BaseFilters) Create(limit uint, offset uint, order BaseOrder) BaseFilters {
	return BaseFilters{Limit: limit, Offset: offset, Order: order}
}
