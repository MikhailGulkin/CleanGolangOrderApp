package command

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/common"
	"github.com/go-faker/faker/v4"
	"testing"
)

type TestEventStore struct {
}

func (t *TestEventStore) Create(_ common.Aggregate, _ interface{}) error {
	return nil
}
func (t *TestEventStore) Update(_ common.Aggregate, _ interface{}) error {
	return nil
}
func (t *TestEventStore) Find(_ string) (common.Aggregate, error) {
	return nil, nil
}

type TestOutbox struct {
}

func (t *TestOutbox) AddEvents(_ []common.Event, _ interface{}) error {
	return nil
}

func TestSuccessCreateCustomerHandle(t *testing.T) {
	createCustomerCommand := CreateCustomerCommand{}
	if err := faker.FakeData(&createCustomerCommand); err != nil {
		t.Fatal(err)
	}
	createCustomer := NewCreateCustomerHandler(&TestEventStore{}, &TestOutbox{})
	customerDTO, err := createCustomer.Handle(createCustomerCommand)
	if err != nil {
		t.Fatal(err)
	}
	if customerDTO.CustomerID != createCustomerCommand.CustomerID.String() {
		t.Fatalf("expected customerID: %s, got: %s", createCustomerCommand.CustomerID, customerDTO.CustomerID)
	}
	if customerDTO.EventID == "" {
		t.Fatalf("expected eventID: %s, got: %s", "", customerDTO.EventID)
	}
}
