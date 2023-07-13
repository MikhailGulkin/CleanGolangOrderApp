package user

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/aggregate/address"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/aggregate/user"
	entity "github.com/MikhailGulkin/simpleGoOrderApp/src/domain/entities/user"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/vo"
)

type Service struct {
}

func (Service) CreateUser(userID vo.UserID, username string, address address.Address) user.User {
	userAddress := entity.UserAddress{}.Create(address.AddressID.Value, address.GetFullAddress())
	userEntity := user.User{}.Create(userID, username, userAddress)
	return userEntity
}
