package address

import (
	"errors"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/address/exceptions"
	appRepo "github.com/MikhailGulkin/simpleGoOrderApp/src/application/address/interfaces/persistence/repo"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/address/aggregate"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/common/id"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/models"
	repo "github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/repo"
	"gorm.io/gorm"
)

type RepoImpl struct {
	repo.BaseGormRepo
	appRepo.AddressRepo
}

func (repo *RepoImpl) AcquireAddressByID(addressID id.ID) (aggregate.Address, error) {
	var addressModel models.Address
	result := repo.Session.Where("id = ?", addressID.ToString()).First(&addressModel)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		exception := exceptions.AddressIDNotExist{}.Exception(addressID.ToString())
		return aggregate.Address{}, &exception
	}
	if result.Error != nil {
		return aggregate.Address{}, result.Error
	}
	return ConvertAddressModelToAggregate(addressModel), nil
}
func (repo *RepoImpl) AddAddress(entity aggregate.Address, tx interface{}) error {
	model := models.Address{
		Base:           models.Base{ID: entity.AddressID.Value},
		BuildingNumber: entity.BuildingNumber,
		StreetName:     entity.StreetName,
		City:           entity.City,
		Country:        entity.Country,
	}
	return tx.(*gorm.DB).Create(&model).Error
}
