package repo

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/common/id"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/user/aggregate"
)

type UserRepo interface {
	AcquireUserByID(userID id.ID) (aggregate.User, error)
	AddUser(user aggregate.User, tx interface{}) error
}
