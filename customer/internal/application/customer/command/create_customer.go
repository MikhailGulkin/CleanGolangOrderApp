package command

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/application/persistence"
	aggregate "github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/aggregate"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/vo"
)

//import (
//	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/application/common/interfaces/persistence"
//	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/application/customer/interfaces/command"
//	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/application/customer/interfaces/persistence/dao"
//	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/application/customer/interfaces/persistence/repo"
//	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/user/services"
//	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/user/vo"
//)
//

type CreateCustomerHandler interface {
	Handle(command CreateCustomerCommand) (CustomerCreateDTO, error)
}

type createCustomerHandler struct {
	persistence.EventStore
	persistence.Outbox
}

func NewCreateCustomerHandler(es persistence.EventStore, o persistence.Outbox) CreateCustomerHandler {
	return &createCustomerHandler{
		EventStore: es,
		Outbox:     o,
	}
}

func (c *createCustomerHandler) Handle(command CreateCustomerCommand) (CustomerCreateDTO, error) {
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
	if err := c.EventStore.Create(customer, nil); err != nil {
		return CustomerCreateDTO{}, err
	}
	if err := c.Outbox.AddEvents(customer.GetUncommittedEvents(), nil); err != nil {
		return CustomerCreateDTO{}, err
	}
	event, eventError := customer.GetLastUncommittedEvent()
	if eventError != nil {
		return CustomerCreateDTO{}, err
	}
	return CustomerCreateDTO{
		CustomerID: customer.Customer.StringID(),
		EventID:    event.GetEventID(),
	}, nil
}
