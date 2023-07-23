package user

import (
	"errors"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/user/exceptions"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/user/interfaces/persistence/dao"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/domain/user/entities"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/db/models"
	repo "github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/db/repo"
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
