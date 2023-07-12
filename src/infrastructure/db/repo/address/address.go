package address

import (
	appRepo "github.com/MikhailGulkin/simpleGoOrderApp/src/application/address/interfaces/persistence/repo"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/aggregate/address"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/models"
	repo "github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/repo"
	"gorm.io/gorm"
)

type RepoImpl struct {
	repo.BaseGormRepo
	appRepo.AddressRepo
}

func (repo *RepoImpl) AddAddress(entity address.Address, tx interface{}) error {
	model := models.Address{
		Base:           models.Base{ID: entity.AddressID.Value},
		BuildingNumber: entity.BuildingNumber,
		StreetName:     entity.StreetName,
		City:           entity.City,
		Country:        entity.Country,
	}
	return tx.(*gorm.DB).Create(&model).Error
}
