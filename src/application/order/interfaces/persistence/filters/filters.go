package filters

type BaseOrder string

const (
	ASC  BaseOrder = "asc"
	DESC BaseOrder = "desc"
)

func ConvertToOrder(s string) BaseOrder {
	if s == string(ASC) {
		return ASC
	}
	if s == string(DESC) {
		return DESC
	}
	return ASC
}

type BaseFilters struct {
	Limit  uint
	Offset uint
	Order  BaseOrder
}

func (filter BaseFilters) Create(limit uint, offset uint, order BaseOrder) BaseFilters {
	return BaseFilters{Limit: limit, Offset: offset, Order: order}
}

type GetAllProductsFilters struct {
	BaseFilters
}
