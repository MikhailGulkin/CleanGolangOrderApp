package order

import "github.com/google/uuid"

type OrderClient struct {
	ClientID uuid.UUID
	Username string
}

func (OrderClient) Create(id uuid.UUID, username string) (OrderClient, error) {
	return OrderClient{
		ClientID: id,
		Username: username,
	}, nil
}
