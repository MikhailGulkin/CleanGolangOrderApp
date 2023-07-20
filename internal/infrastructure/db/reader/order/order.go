package order

import (
	"errors"
	"fmt"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/dto"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/exceptions"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/interfaces/persistence/filters"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/interfaces/persistence/reader"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/db/models"
	r "github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/db/reader"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReaderImpl struct {
	reader.OrderReader
	r.BaseGormDAO
}

func (dao *ReaderImpl) GetAllOrders(filters filters.GetAllOrdersFilters) ([]dto.Order, error) {
	var orders []models.Order
	result := dao.Session.
		Preload("Address").
		Preload("Client").
		Preload("Products").
		Limit(int(filters.Limit)).
		Offset(int(filters.Offset)).
		Order(fmt.Sprintf("created_at %s", filters.Order)).
		Where("closed = ? AND saga_status = ?", false, "Approved").
		Find(&orders)
	if result.Error != nil {
		return []dto.Order{}, result.Error
	}
	return ConvertOrderModelsToDTOs(orders), nil
}
func (dao *ReaderImpl) GetOrderByID(id uuid.UUID) (dto.Order, error) {
	var order models.Order
	result := dao.Session.
		Preload("Address").
		Preload("Client").
		Preload("Products").
		Where("closed = ?", false).
		Where("id = ?", id).
		First(&order)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		exception := exceptions.OrderIDNotExist{}.Exception(id.String())
		return dto.Order{}, &exception
	}
	return ConvertOrderModelToDTO(order), nil
}
