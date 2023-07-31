package repo

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/user/aggregate"
)

type UserRepo interface {
	AddUser(user aggregate.User, tx interface{}) error
}