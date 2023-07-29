package aggregate

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/user/entities"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/user/vo"
)

type User struct {
	vo.UserID
	Username string
	Address  entities.UserAddress
	Orders   []entities.UserOrder
}

func (User) Create(id vo.UserID, username string, address entities.UserAddress) User {
	return User{
		UserID:   id,
		Username: username,
		Address:  address,
	}
}
