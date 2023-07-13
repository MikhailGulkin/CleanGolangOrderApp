package repo

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/aggregate/user"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/common/id"
)

type UserRepo interface {
	AcquireUserByID(userID id.ID) (user.User, error)
	AddUser(user user.User, tx interface{}) error
}
