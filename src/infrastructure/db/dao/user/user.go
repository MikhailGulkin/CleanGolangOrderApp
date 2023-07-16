package user

import (
	"errors"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/user/exceptions"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/user/interfaces/persistence/dao"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/user/entities"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/models"
	repo "github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/repo"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DAOImpl struct {
	dao.UserDAO
	repo.BaseGormRepo
}

func (dao *DAOImpl) GetUserAddress(addressID uuid.UUID) (entities.UserAddress, error) {
	var addressModel models.Address
	result := dao.Session.Where("id = ?", addressID.String()).First(&addressModel)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		exception := exceptions.UserAddressIDNotExist{}.Exception(addressID.String())
		return entities.UserAddress{}, &exception
	}
	if result.Error != nil {
		return entities.UserAddress{}, result.Error
	}
	return ConvertModelToUserAddressEntities(addressModel), nil
}
