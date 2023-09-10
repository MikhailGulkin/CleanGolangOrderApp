package aggregate

import (
	"fmt"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/common"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/entities"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/vo"
	"github.com/google/uuid"
)

const (
	CustomerAggregateType common.AggregateType = "customer"
)

type Customer struct {
	FullName     vo.FullName
	AddressID    uuid.UUID
	Email        vo.Email
	AvatarUri    string
	Balance      vo.Balance
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
	return fmt.Sprintf(
		"CustomerAggregate{ID: %s, FullName: %s, AddressID: %s, Email: %s, Balance: %s, Transactions: %v, Orders: %v, AvatarUri: %s}",
		a.GetID(), a.Customer.FullName, a.Customer.AddressID, a.Customer.Email, a.Customer.Balance, a.Customer.Transactions, a.Customer.Orders, a.Customer.AvatarUri)
}

func NewCustomer() *Customer {
	return &Customer{
		Transactions: make([]*entities.CustomerTransactions, 0),
		Orders:       make([]*uuid.UUID, 0),
	}
}

func NewCustomerAggregateWithID(id uuid.UUID) *CustomerAggregate {
	if id.String() == "" {
		return nil
	}

	aggregate := NewCustomerAggregate()
	aggregate.SetID(id.String())
	return aggregate
}

func NewCustomerAggregate() *CustomerAggregate {
	orderAggregate := &CustomerAggregate{Customer: NewCustomer()}
	base := common.NewAggregateBase(orderAggregate.When)
	base.SetType(CustomerAggregateType)
	orderAggregate.AggregateBase = base
	return orderAggregate
}
