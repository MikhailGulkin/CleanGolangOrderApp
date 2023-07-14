package address

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/address/aggregate"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/address/vo"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/models"
)

func ConvertAddressModelToAggregate(model models.Address) aggregate.Address {
	return aggregate.Address{
		AddressID:      vo.AddressID{Value: model.ID},
		BuildingNumber: model.BuildingNumber,
		StreetName:     model.StreetName,
		City:           model.City,
		Country:        model.Country,
	}
}
