package persistence

import (
	"context"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/common"
)

type Bucket interface {
	UploadAvatar(ctx context.Context, name string, fileBuffer []byte) error
}
type Outbox interface {
	AddEvents(events []common.Event, tx interface{}) error
}
type UoW interface {
	Commit() error
	Rollback() error
	Begin() (interface{}, error)
}
type UoWManager interface {
	GetUoW() UoW
}
