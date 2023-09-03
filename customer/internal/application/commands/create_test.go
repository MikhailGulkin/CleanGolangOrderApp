package commands

import (
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/mock"
	"strings"
	"testing"
)

func TestSuccessCreateCustomerHandle(t *testing.T) {
	createCustomerCommand := CreateCustomerCommand{}
	if err := faker.FakeData(&createCustomerCommand); err != nil {
		t.Fatal(err)
	}
	eventStore := new(MockEventStore)
	eventStore.On("Exists", mock.Anything).Return(nil)
	eventStore.On("Create", mock.Anything, mock.Anything).Return(nil)

	createCustomerCommand.FirstName = "Nike"
	createCustomerCommand.MiddleName = "Vladimirovich"
	createCustomerCommand.LastName = "Sulkin"

	createCustomer := NewCreateCustomerHandler(eventStore, &TestOutbox{}, &TestUoWManager{})

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
