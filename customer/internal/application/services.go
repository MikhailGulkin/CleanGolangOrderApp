package application

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/application/commands"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/application/persistence"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/common"
)

type CustomerServices struct {
	Commands *commands.CustomerCommands
}
type TestOutbox struct {
}

func (t *TestOutbox) AddEvents(_ []common.Event, _ interface{}) error {
	return nil
}
func NewCustomerServices(es persistence.EventStore, manager persistence.UoWManager) *CustomerServices {
	createCustomerHandler := commands.NewCreateCustomerHandler(es, &TestOutbox{}, manager)

	customerCommands := commands.NewCustomerCommands(createCustomerHandler)
	return &CustomerServices{
		Commands: customerCommands,
	}
}
