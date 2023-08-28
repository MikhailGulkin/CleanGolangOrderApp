package commands

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/application/persistence"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/common"
	"github.com/go-faker/faker/v4"
	"strings"
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
func (t *TestEventStore) Exists(_ string) error {
	return nil
}

type TestOutbox struct {
}

func (t *TestOutbox) AddEvents(_ []common.Event, _ interface{}) error {
	return nil
}

type TestUoW struct {
}

func (t *TestUoW) Begin() (interface{}, error) {
	return nil, nil
}
func (t *TestUoW) Commit() error {
	return nil
}
func (t *TestUoW) Rollback() error {
	return nil
}

type TestUoWManager struct {
}

func (t *TestUoWManager) GetUoW() persistence.UoW {
	return &TestUoW{}
}

func TestSuccessCreateCustomerHandle(t *testing.T) {
	createCustomerCommand := CreateCustomerCommand{}
	if err := faker.FakeData(&createCustomerCommand); err != nil {
		t.Fatal(err)
	}
	createCustomer := NewCreateCustomerHandler(&TestEventStore{}, &TestOutbox{}, &TestUoWManager{})
	customerDTO, err := createCustomer.Handle(createCustomerCommand)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(customerDTO.CustomerID, createCustomerCommand.CustomerID.String()) == false {
		t.Fatalf("expected customerID: %s, got: %s", createCustomerCommand.CustomerID, customerDTO.CustomerID)
	}
	if customerDTO.EventID == "" {
		t.Fatalf("expected eventID: %s, got: %s", "", customerDTO.EventID)
	}
}
