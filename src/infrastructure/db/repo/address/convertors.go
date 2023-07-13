package address

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/aggregate/address"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/vo"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/models"
)

func ConvertAddressModelToAggregate(model models.Address) address.Address {
	return address.Address{
		AddressID:      vo.AddressID{Value: model.ID},
		BuildingNumber: model.BuildingNumber,
		StreetName:     model.StreetName,
		City:           model.City,
		Country:        model.Country,
	}
}
