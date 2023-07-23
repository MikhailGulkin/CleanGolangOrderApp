package user

import (
	appRepo "github.com/MikhailGulkin/simpleGoOrderApp/internal/application/user/interfaces/persistence/repo"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/domain/user/aggregate"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/db/models"
	repo "github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/db/repo"
	"gorm.io/gorm"
)

type RepoImpl struct {
	repo.BaseGormRepo
	appRepo.UserRepo
}

func (repo *RepoImpl) AddUser(entity aggregate.User, tx any) error {
	model := models.User{
		Base:      models.Base{ID: entity.UserID.Value},
		Username:  entity.Username,
		AddressID: entity.Address.AddressID,
	}
	return tx.(*gorm.DB).Create(&model).Error
}
