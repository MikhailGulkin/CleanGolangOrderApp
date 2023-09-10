package application

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/application/commands"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/application/persistence"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/aggregate"
)

type CustomerServices struct {
	Commands *commands.CustomerCommands
}

func NewCustomerServices(
	es aggregate.EventStore,
	outbox persistence.Outbox,
	manager persistence.UoWManager,
	bucket persistence.Bucket,
) *CustomerServices {
	createCustomerHandler := commands.NewCreateCustomerHandler(es, outbox, manager)
	uploadCustomerAvatarHandler := commands.NewCustomerUploadAvatarHandler(es, outbox, manager, bucket)

	customerCommands := commands.NewCustomerCommands(createCustomerHandler, uploadCustomerAvatarHandler)
	return &CustomerServices{
		Commands: customerCommands,
	}
}
