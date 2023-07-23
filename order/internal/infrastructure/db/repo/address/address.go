package address

import (
	appRepo "github.com/MikhailGulkin/simpleGoOrderApp/internal/application/address/interfaces/persistence/repo"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/domain/address/aggregate"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/db/models"
	repo "github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/db/repo"
	"gorm.io/gorm"
)

type RepoImpl struct {
	repo.BaseGormRepo
	appRepo.AddressRepo
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
