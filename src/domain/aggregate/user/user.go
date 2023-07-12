package user

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/entities/user"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/vo"
)

type User struct {
	vo.UserID
	Username string
	Orders   []user.UserOrder
}
