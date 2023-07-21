package cache

import (
	"github.com/google/uuid"
	"time"
)

type EventType struct {
	EventType string `json:"eventType"`
}
type OrderCreateClientSubscribe struct {
	ClientID uuid.UUID `json:"clientID"`
	Username string    `json:"username,omitempty"`
}
type OrderCreatedAddressSubscribe struct {
	AddressID   uuid.UUID `json:"clientID"`
	FullAddress string    `json:"username"`
}
type OrderCreateSubscribe struct {
	OrderID         uuid.UUID                    `json:"orderID"`
	Client          OrderCreateClientSubscribe   `json:"client"`
	PaymentMethod   string                       `json:"paymentMethod"`
	OrderStatus     string                       `json:"orderStatus"`
	DeliveryAddress OrderCreatedAddressSubscribe `json:"deliveryAddress"`
	SerialNumber    int                          `json:"serialNumber"`
	TotalPrice      float64                      `json:"totalPrice"`
	CreatedAt       time.Time                    `json:"eventTimeStamp"`
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

type OrderCache interface {
	OrderEvent(event interface{})
}
