package events

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/domain/common/events"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/domain/order/consts"
	"github.com/google/uuid"
)

type OrderCreatedClient struct {
	ClientID uuid.UUID `json:"clientID"`
	Username string    `json:"username"`
}
type OrderCreatedAddress struct {
	AddressID   uuid.UUID `json:"clientID"`
	FullAddress string    `json:"username"`
}
type OrderCreated struct {
	events.BaseEvent
	OrderID         uuid.UUID           `json:"orderID"`
	Client          OrderCreatedClient  `json:"client"`
	OrderStatus     string              `json:"orderStatus"`
	PaymentMethod   string              `json:"paymentMethod"`
	DeliveryAddress OrderCreatedAddress `json:"deliveryAddress"`
	SerialNumber    int                 `json:"serialNumber"`
	TotalPrice      float64             `json:"totalPrice"`
}

func (OrderCreated) Create(
	orderID uuid.UUID,
	client OrderCreatedClient,
	paymentMethod string,
	serialNumber int,
	totalPrice float64,
	address OrderCreatedAddress,
) events.Event {
	return &OrderCreated{
		OrderID:         orderID,
		BaseEvent:       events.BaseEvent{}.Create("OrderCreated"),
		Client:          client,
		PaymentMethod:   paymentMethod,
		DeliveryAddress: address,
		OrderStatus:     string(consts.New),
		SerialNumber:    serialNumber,
		TotalPrice:      totalPrice,
	}
}

func (o *OrderCreated) Bytes() ([]byte, error) {
	return events.Bytes(o)
}
func (o *OrderCreated) UniqueAggregateID() uuid.UUID {
	return o.OrderID
}
