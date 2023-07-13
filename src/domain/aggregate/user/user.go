package user

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/entities/user"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/vo"
)

type User struct {
	vo.UserID
	Username string
	Address  user.UserAddress
	Orders   []user.UserOrder
}

func (User) Create(id vo.UserID, username string, address user.UserAddress) User {
	return User{
		UserID:   id,
		Username: username,
		Address:  address,
	}
}
