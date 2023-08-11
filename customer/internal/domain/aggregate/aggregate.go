package aggregate

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/common"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/entities"
	vo2 "github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/vo"
	"github.com/google/uuid"
)

const (
	OrderAggregateType common.AggregateType = "order"
)

type Customer struct {
	vo2.UserID
	FullName     vo2.FullName
	AddressID    uuid.UUID
	Transactions []*entities.CustomerTransactions
	Balance      vo2.CustomerBalance
	Orders       []*uuid.UUID
}
type CustomerAggregate struct {
	Customer *Customer
	*common.AggregateBase
}

func NewCustomer() *Customer {
	return &Customer{
		Transactions: make([]*entities.CustomerTransactions, 0),
		Orders:       make([]*uuid.UUID, 0),
	}
}

func NewOrderAggregateWithID(id string) *CustomerAggregate {
	if id == "" {
		return nil
	}

	aggregate := NewOrderAggregate()
	aggregate.SetID(id)
	aggregate.Customer.SetID(id)
	return aggregate
}

func NewOrderAggregate() *CustomerAggregate {
	orderAggregate := &CustomerAggregate{Customer: NewCustomer()}
	base := common.NewAggregateBase(orderAggregate.When)
	base.SetType(OrderAggregateType)
	orderAggregate.AggregateBase = base
	return orderAggregate
}
