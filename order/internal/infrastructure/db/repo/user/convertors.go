package user

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/domain/user/aggregate"
	entity "github.com/MikhailGulkin/simpleGoOrderApp/order/internal/domain/user/entities"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/domain/user/vo"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/db/models"
)

func ConvertUserModelToAggregate(model models.User) aggregate.User {
	return aggregate.User{
		UserID:   vo.UserID{Value: model.ID},
		Username: model.Username,
		Address: entity.UserAddress{
			AddressID: model.Address.ID,
			Address:   model.Address.GetFullAddress(),
		},
	}
}
