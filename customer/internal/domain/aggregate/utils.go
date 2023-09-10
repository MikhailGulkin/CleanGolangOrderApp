package aggregate

import (
	"context"
	"fmt"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/common"
	"github.com/google/uuid"
)

func LoadOrderAggregate(_ context.Context, eventStore EventStore, aggregateID uuid.UUID) (*CustomerAggregate, error) {
	customer := NewCustomerAggregateWithID(aggregateID)

	err := eventStore.Exists(customer.GetID())
	if err == nil {
		return nil, common.ErrAggregateNotFound
	}
	if err := eventStore.Load(customer); err != nil {
		return nil, err
	}
	fmt.Printf("Aggregate loaded: %v\n", customer)
	return customer, nil
}
