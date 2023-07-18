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
