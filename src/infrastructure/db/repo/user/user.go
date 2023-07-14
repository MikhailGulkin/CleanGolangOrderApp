package user

import (
	"errors"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/user/exceptions"
	appRepo "github.com/MikhailGulkin/simpleGoOrderApp/src/application/user/interfaces/persistence/repo"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/common/id"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/user/aggregate"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/models"
	repo "github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/repo"
	"gorm.io/gorm"
)

type RepoImpl struct {
	repo.BaseGormRepo
	appRepo.UserRepo
}

func (repo *RepoImpl) AcquireUserByID(userID id.ID) (aggregate.User, error) {
	var userModel models.User
	result := repo.Session.Preload("Address").Preload("Orders").Where("id = ?", userID.ToString()).First(&userModel)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		exception := exceptions.UserIDNotExist{}.Exception(userID.ToString())
		return aggregate.User{}, &exception
	}
	if result.Error != nil {
		return aggregate.User{}, result.Error
	}
	return ConvertUserModelToAggregate(userModel), nil
}

func (repo *RepoImpl) AddUser(entity aggregate.User, tx any) error {
	model := models.User{
		Base:      models.Base{ID: entity.UserID.Value},
		Username:  entity.Username,
		AddressID: entity.Address.AddressID,
	}
	return tx.(*gorm.DB).Create(&model).Error
}
