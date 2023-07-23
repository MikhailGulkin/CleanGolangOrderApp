package cache

import (
	"github.com/google/uuid"
	"time"
)

type EventType struct {
	EventType string `json:"eventType"`
}

type OrderCreateSubscribe struct {
	OrderID           uuid.UUID `json:"orderID"`
	ClientID          uuid.UUID `json:"clientID"`
	PaymentMethod     string    `json:"paymentMethod"`
	OrderStatus       string    `json:"orderStatus"`
	DeliveryAddressID uuid.UUID `json:"deliveryAddressID"`
	SerialNumber      int       `json:"serialNumber"`
	TotalPrice        float64   `json:"totalPrice"`
	CreatedAt         time.Time `json:"eventTimeStamp"`
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
type OrderDeleteEvent struct {
	OrderID uuid.UUID `json:"orderID"`
}
type OrderCache interface {
	OrderCreateEvent(event OrderCreateSubscribe)
	OrderAddProductEvent(event OrderAddProductSubscribe)
	OrderDeleteEvent(event OrderDeleteEvent)
}
