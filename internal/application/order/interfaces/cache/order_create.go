package cache

import (
	"github.com/google/uuid"
	"time"
)

type OrderCreateClientSubscribe struct {
	ClientID uuid.UUID `json:"clientID,omitempty"`
	Username string    `json:"username,omitempty"`
}
type OrderCreateSubscribe struct {
	OrderID         uuid.UUID                  `json:"orderID,omitempty"`
	Client          OrderCreateClientSubscribe `json:"client"`
	PaymentMethod   string                     `json:"paymentMethod,omitempty"`
	OrderStatus     string                     `json:"orderStatus"`
	DeliveryAddress uuid.UUID                  `json:"deliveryAddress,omitempty"`
	SerialNumber    int                        `json:"serialNumber,omitempty"`
	TotalPrice      float64                    `json:"totalPrice,omitempty"`
}
type OrderAddProductSubscribe struct {
	OrderID      uuid.UUID `json:"orderID"`
	ClientID     uuid.UUID `json:"clientID"`
	ProductName  string    `json:"productName"`
	ProductID    uuid.UUID `json:"productID"`
	ProductPrice float64   `json:"productPrice"`
}
type ProductEvent struct {
	ProductID   uuid.UUID `json:"productID"`
	Name        string    `json:"name"`
	ActualPrice float64   `json:"actualPrice"`
}
type ClientEvent struct {
	ClientID   uuid.UUID
	ClientName string
}

type OrderEvent struct {
	OrderStatus   string         `json:"orderStatus"`
	PaymentMethod string         `json:"paymentMethod"`
	Client        ClientEvent    `json:"client"`
	Address       string         `json:"address"`
	OrderID       uuid.UUID      `json:"OrderID"`
	SerialNumber  int            `json:"serialNumber"`
	Products      []ProductEvent `json:"products"`
	CreatedAt     time.Time      `json:"createdAt"`
	TotalPrice    float64        `json:"totalPrice"`
}
type OrderCache interface {
	OrderCreate(event interface{})
}
