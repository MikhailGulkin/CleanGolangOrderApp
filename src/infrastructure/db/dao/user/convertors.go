package user

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/user/entities"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/models"
)

func ConvertModelToUserAddressEntities(model models.Address) entities.UserAddress {
	return entities.UserAddress{
		AddressID: model.ID,
		Address:   model.GetFullAddress(),
	}
}
