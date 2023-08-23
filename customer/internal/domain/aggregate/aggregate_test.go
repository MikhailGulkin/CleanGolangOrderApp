package aggregate

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/consts"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/events"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/vo"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"testing"
)

func TestCustomer(t *testing.T) {
	id := uuid.New()
	customer := NewCustomerAggregateWithID(id)
	if customer.Customer.CustomerID.Value != id {
		t.Fatal("wrong id")
	}
	t.Run("create customer", CreateCustomerTest(customer))
	t.Run("update transaction", TransactionFreezeUpdateEvent(customer))
	t.Run("update transaction", TransactionDepositUpdateEvent(customer))

}
func CreateCustomerTest(customer *CustomerAggregate) func(t *testing.T) {
	return func(t *testing.T) {
		createEvent := events.CustomerCreatedEvent{}
		err := faker.FakeData(&createEvent)
		if err != nil {
			t.Fatal(err)
		}

		if err := customer.CreateCustomer(createEvent.FullName, createEvent.AddressID, createEvent.Email); err != nil {
			t.Fatal(err)
		}
		if customer.Customer.Email != createEvent.Email {
			t.Fatal("wrong email")
		}
		if customer.Customer.FullName != createEvent.FullName {
			t.Fatal("wrong full name")
		}
		if customer.Customer.AddressID != createEvent.AddressID {
			t.Fatal("wrong address id")
		}
		if customer.Customer.Balance != vo.NewBalance() {
			t.Fatal("wrong balance")
		}
		if len(customer.GetUncommittedEvents()) != 1 {
			t.Fatal("wrong number of uncommitted events")
		}
	}

}
func TransactionFreezeUpdateEvent(customer *CustomerAggregate) func(t *testing.T) {
	return func(t *testing.T) {
		transactionEvent := events.TransactionsUpdatedEvent{}
		err := faker.FakeData(&transactionEvent)
		if err != nil {
			t.Fatal(err)
		}
		transactionEvent.Transaction.TransactionType = consts.PURCHASE_PENDING
		err = customer.UpdateTransactionCustomer(transactionEvent.Transaction)
		if err != nil {
			t.Fatal(err)
		}
		if len(customer.Customer.Transactions) != 1 {
			t.Fatal("wrong number of transactions")
		}

		if len(customer.GetUncommittedEvents()) != 3 {
			t.Fatal("wrong number of uncommitted events")
		}
		if (customer.Customer.Balance.AvailableMoney.Value + transactionEvent.Transaction.TransactionSum.Value) != 0 {
			t.Fatal("wrong balance")
		}
		if (customer.Customer.Balance.FrozenMoney.Value - transactionEvent.Transaction.TransactionSum.Value) != 0 {
			t.Fatal("wrong balance")
		}
	}
}
func TransactionDepositUpdateEvent(customer *CustomerAggregate) func(t *testing.T) {
	return func(t *testing.T) {
		transactionEvent := events.TransactionsUpdatedEvent{}
		err := faker.FakeData(&transactionEvent)
		if err != nil {
			t.Fatal(err)
		}
		transactionEvent.Transaction.TransactionType = consts.DEPOSIT
		err = customer.UpdateTransactionCustomer(transactionEvent.Transaction)
		if err != nil {
			t.Fatal(err)
		}
		if len(customer.Customer.Transactions) != 2 {
			t.Fatal("wrong number of transactions")
		}
		if len(customer.GetUncommittedEvents()) != 5 {
			t.Fatal("wrong number of uncommitted events")
		}
	}
}
