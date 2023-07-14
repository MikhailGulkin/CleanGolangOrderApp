package services

import (
	aggregate2 "github.com/MikhailGulkin/simpleGoOrderApp/src/domain/address/aggregate"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/user/aggregate"
	entity "github.com/MikhailGulkin/simpleGoOrderApp/src/domain/user/entities"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/user/vo"
)

type Service struct {
}

func (Service) CreateUser(userID vo.UserID, username string, address aggregate2.Address) aggregate.User {
	userAddress := entity.UserAddress{}.Create(address.AddressID.Value, address.GetFullAddress())
	userEntity := aggregate.User{}.Create(userID, username, userAddress)
	return userEntity
}
