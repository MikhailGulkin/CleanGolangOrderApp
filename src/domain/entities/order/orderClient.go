package order

type OrderClient struct {
}

func (OrderClient) Create(_ string) (OrderClient, error) {
	return OrderClient{}, nil
}
