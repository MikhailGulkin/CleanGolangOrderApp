package graphql

import (
	"context"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer-gateway/internal/grpc/pb"
	"time"
)

type mutationResolver struct {
	client pb.CustomerServiceClient
}

func (r *mutationResolver) CreateCustomer(ctx context.Context, input CustomerInput) (*CustomerResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	customer, err := r.client.CreateCustomer(ctx, &pb.CreateCustomerRequest{
		CustomerID: input.ID,
		FirstName:  input.FirstName,
		MiddleName: input.MiddleName,
		LastName:   input.LastName,
		Email:      input.Email,
		AddressID:  input.AddressID,
	})
	if err != nil {
		return nil, err
	}
	return &CustomerResponse{
		customer.Id,
		customer.EventID,
	}, nil
}
