package commands

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/application/persistence"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/aggregate"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/vo"
)

type CreateCustomerHandler struct {
	aggregate.EventStore
	persistence.Outbox
	persistence.UoWManager
}

func NewCreateCustomerHandler(
	es aggregate.EventStore,
	o persistence.Outbox,
	m persistence.UoWManager,
) *CreateCustomerHandler {
	return &CreateCustomerHandler{
		EventStore: es,
		Outbox:     o,
		UoWManager: m,
	}
}

func (c *CreateCustomerHandler) Handle(command CreateCustomerCommand) (CustomerCreateDTO, error) {
	fullName, err := vo.NewFullName(command.FirstName, command.MiddleName, command.LastName)
	if err != nil {
		return CustomerCreateDTO{}, err
	}
	email, errEmail := vo.NewEmail(command.Email)
	if errEmail != nil {
		return CustomerCreateDTO{}, errEmail
	}
	customer := aggregate.NewCustomerAggregateWithID(command.CustomerID)
	if err := customer.CreateCustomer(fullName, command.AddressID, email); err != nil {
		return CustomerCreateDTO{}, err
	}

	if err := c.EventStore.Exists(customer.GetID()); err != nil {
		return CustomerCreateDTO{}, err
	}
	uow := c.UoWManager.GetUoW()

	tx, errTx := uow.Begin()
	if errTx != nil {
		return CustomerCreateDTO{}, errTx
	}
	if err := c.EventStore.Create(customer, tx); err != nil {
		return CustomerCreateDTO{}, err
	}
	if err := c.Outbox.AddEvents(customer.GetUncommittedEvents(), tx); err != nil {
		return CustomerCreateDTO{}, err
	}
	event, eventError := customer.GetLastUncommittedEvent()
	if eventError != nil {
		err := uow.Rollback()
		if err != nil {
			return CustomerCreateDTO{}, err
		}
		return CustomerCreateDTO{}, eventError
	}
	if err := uow.Commit(); err != nil {
		return CustomerCreateDTO{}, err
	}
	return CustomerCreateDTO{
		CustomerID: customer.GetID(),
		EventID:    event.GetEventID(),
	}, nil
}
