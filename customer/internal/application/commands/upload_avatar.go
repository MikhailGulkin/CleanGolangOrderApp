package commands

import (
	"context"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/application/persistence"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/aggregate"
)

type UploadAvatarCustomerHandler interface {
	Handle(command UploadAvatarCommand) (UploadAvatarDTO, error)
}

type uploadCustomerHandler struct {
	aggregate.EventStore
	persistence.Outbox
	persistence.UoWManager
	persistence.Bucket
}

func NewCustomerUploadAvatarHandler(
	es aggregate.EventStore,
	o persistence.Outbox,
	m persistence.UoWManager,
	bucket persistence.Bucket,
) UploadAvatarCustomerHandler {
	return &uploadCustomerHandler{
		EventStore: es,
		Outbox:     o,
		UoWManager: m,
		Bucket:     bucket,
	}
}
func (c *uploadCustomerHandler) Handle(command UploadAvatarCommand) (UploadAvatarDTO, error) {
	customer, err := aggregate.LoadOrderAggregate(context.Background(), c.EventStore, command.CustomerID)
	if err != nil {
		return UploadAvatarDTO{}, err
	}
	if err := customer.CreateAvatarUri(); err != nil {
		return UploadAvatarDTO{}, err
	}
	err = c.Bucket.UploadAvatar(context.Background(), customer.GetAvatarUri(), command.Avatar)
	if err != nil {
		return UploadAvatarDTO{}, err
	}
	uow := c.UoWManager.GetUoW()
	tx, err := uow.Begin()
	if err != nil {
		return UploadAvatarDTO{}, err
	}

	if err := c.EventStore.Update(customer, tx); err != nil {
		return UploadAvatarDTO{}, err
	}
	if err := c.Outbox.AddEvents(customer.GetUncommittedEvents(), tx); err != nil {
		return UploadAvatarDTO{}, err
	}
	event, eventError := customer.GetLastUncommittedEvent()
	if eventError != nil {
		err := uow.Rollback()
		if err != nil {
			return UploadAvatarDTO{}, err
		}
		return UploadAvatarDTO{}, eventError
	}
	if err := uow.Commit(); err != nil {
		return UploadAvatarDTO{}, err
	}
	return UploadAvatarDTO{
		CustomerID: customer.GetID(),
		AvatarUri:  customer.GetAvatarUri(),
		EventID:    event.GetEventID(),
	}, nil
}
