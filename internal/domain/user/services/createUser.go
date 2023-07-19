package services

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/domain/user/aggregate"
	entity "github.com/MikhailGulkin/simpleGoOrderApp/internal/domain/user/entities"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/domain/user/vo"
)

type Service struct {
}

func (Service) CreateUser(userID vo.UserID, username string, address entity.UserAddress) aggregate.User {
	userEntity := aggregate.User{}.Create(userID, username, address)
	return userEntity
}