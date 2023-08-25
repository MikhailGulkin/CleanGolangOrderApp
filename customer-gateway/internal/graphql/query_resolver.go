package graphql

import "context"

type queryResolver struct {
}

func (q *queryResolver) GetCustomer(ctx context.Context, id SomeQuery) (*CustomerResponse, error) {
	panic("not implemented")
}
