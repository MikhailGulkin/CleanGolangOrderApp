package repo

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/aggregate/user"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/vo"
)

type UserRepo interface {
	AcquireUserByID(userID vo.UserID) (user.User, error)
}
