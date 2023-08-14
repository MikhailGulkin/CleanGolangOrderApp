package aggregate

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/common"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/entities"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/vo"
	"github.com/google/uuid"
)

const (
	OrderAggregateType common.AggregateType = "order"
)

type Customer struct {
	vo.CustomerID
	FullName     vo.FullName
	AddressID    uuid.UUID
	Balance      vo.CustomerBalance
	Transactions []*entities.CustomerTransactions
	Orders       []*uuid.UUID
}
type CustomerAggregate struct {
	Customer *Customer
	*common.AggregateBase
}

func (a *CustomerAggregate) RaiseEvent(event common.Event) error {
	//TODO implement me
	panic("implement me")
}

func (a *CustomerAggregate) String() string {
	//TODO implement me
	panic("implement me")
}

func NewCustomer() *Customer {
	return &Customer{
		Transactions: make([]*entities.CustomerTransactions, 0),
		Orders:       make([]*uuid.UUID, 0),
	}
}

func NewCustomerAggregateWithID(id string) *CustomerAggregate {
	if id == "" {
		return nil
	}

	aggregate := NewCustomerAggregate()
	aggregate.SetID(id)
	aggregate.Customer.SetID(id)
	return aggregate
}

func NewCustomerAggregate() *CustomerAggregate {
	orderAggregate := &CustomerAggregate{Customer: NewCustomer()}
	base := common.NewAggregateBase(orderAggregate.When)
	base.SetType(OrderAggregateType)
	orderAggregate.AggregateBase = base
	return orderAggregate
}
