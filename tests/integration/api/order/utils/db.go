package utils

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/order/aggregate"
	order "github.com/MikhailGulkin/simpleGoOrderApp/src/domain/order/entities"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/order/vo"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateValidOrderEntity() aggregate.Order {
	return aggregate.Order{
		OrderID: vo.OrderID{Value: uuid.New()},
		Client: order.OrderClient{
			ClientID: uuid.New(),
		},
		DeliveryAddress: order.OrderAddress{
			AddressID: uuid.New(),
		},
	}
}
func CreateValidOrderModel(entity aggregate.Order) models.Order {
	return models.Order{
		Base:      models.Base{ID: entity.OrderID.Value},
		Address:   models.Address{Base: models.Base{ID: entity.DeliveryAddress.AddressID}},
		AddressID: entity.DeliveryAddress.AddressID,
		ClientID:  entity.Client.ClientID,
		Client:    models.User{Base: models.Base{ID: entity.Client.ClientID}, AddressID: entity.DeliveryAddress.AddressID},
		Products:  []models.Product{{Base: models.Base{ID: uuid.New()}}},
	}
}
func CreateProductInDB(model *models.Order, db *gorm.DB) {
	db.Create(&model.Products[0])
}
func CreateClientInDB(model *models.Order, db *gorm.DB) {
	db.Create(&model.Client)
}
func CreateAddressInDB(model *models.Order, db *gorm.DB) {
	db.Create(&model.Address)
}
func CreateOrderInDB(model *models.Order, db *gorm.DB) {
	CreateProductInDB(model, db)
	CreateAddressInDB(model, db)
	CreateClientInDB(model, db)
	db.Create(&model)
}
