package dto

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/common/dto"
	"github.com/google/uuid"
	"time"
)

type Order struct {
	OrderStatus   string    `json:"orderStatus"`
	PaymentMethod string    `json:"paymentMethod"`
	ClientID      uuid.UUID `json:"clientID"`
	AddressID     uuid.UUID `json:"addressID"`
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
