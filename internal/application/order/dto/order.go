package dto

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/common/dto"
	"github.com/google/uuid"
	"time"
)

type Order struct {
	OrderStatus   string    `json:"orderStatus"`
	PaymentMethod string    `json:"paymentMethod"`
	Client        Client    `json:"client"`
	Address       string    `json:"address"`
	OrderID       uuid.UUID `json:"OrderID"`
	SerialNumber  int       `json:"serialNumber"`
	Products      []Product `json:"products"`
	CreatedAt     time.Time `json:"createdAt"`
	TotalPrice    float64   `json:"totalPrice"`
}
type Orders struct {
	Orders []Order `json:"orders"`
	dto.BaseSequence
}
type OrderCreateClientSubscribe struct {
	ClientID uuid.UUID `json:"clientID,omitempty"`
	Username string    `json:"username,omitempty"`
}
type OrderCreateSubscribe struct {
	OrderID         uuid.UUID                  `json:"orderID,omitempty"`
	Client          OrderCreateClientSubscribe `json:"client"`
	PaymentMethod   string                     `json:"paymentMethod,omitempty"`
	DeliveryAddress uuid.UUID                  `json:"deliveryAddress,omitempty"`
	SerialNumber    int                        `json:"serialNumber,omitempty"`
	TotalPrice      float64                    `json:"totalPrice,omitempty"`
}
