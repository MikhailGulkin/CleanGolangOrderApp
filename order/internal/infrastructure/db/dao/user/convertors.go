package user

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/domain/user/entities"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/db/models"
)

func ConvertModelToUserAddressEntities(model models.Address) entities.UserAddress {
	return entities.UserAddress{
		AddressID: model.ID,
		Address:   model.GetFullAddress(),
	}
}
