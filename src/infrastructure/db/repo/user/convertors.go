package user

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/aggregate/user"
	entity "github.com/MikhailGulkin/simpleGoOrderApp/src/domain/entities/user"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/vo"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/models"
)

func ConvertUserModelToAggregate(model models.User) user.User {
	return user.User{
		UserID:   vo.UserID{Value: model.ID},
		Username: model.Username,
		Address: entity.UserAddress{
			AddressID: model.Address.ID,
			Address:   model.Address.GetFullAddress(),
		},
	}
}
