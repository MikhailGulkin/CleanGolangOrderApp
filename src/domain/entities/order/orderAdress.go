package order

type OrderAddress struct {
}

func (OrderAddress) Create(_ int) (OrderAddress, error) {
	return OrderAddress{}, nil
}
