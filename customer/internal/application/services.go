package application

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/application/commands"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/application/persistence"
)

type CustomerServices struct {
	Commands *commands.CustomerCommands
}

func NewCustomerServices(es persistence.EventStore, outbox persistence.Outbox, manager persistence.UoWManager) *CustomerServices {
	createCustomerHandler := commands.NewCreateCustomerHandler(es, outbox, manager)

	customerCommands := commands.NewCustomerCommands(createCustomerHandler)
	return &CustomerServices{
		Commands: customerCommands,
	}
}
