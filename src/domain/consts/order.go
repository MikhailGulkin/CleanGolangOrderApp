package consts

type OrderStatus string

const (
	New          OrderStatus = "New"
	InProcessing OrderStatus = "InProcessing"
	Processed    OrderStatus = "Processed"
	Delivered    OrderStatus = "Delivered"
	Canceled     OrderStatus = "Canceled"
)

type PaymentMethod string

const (
	Online PaymentMethod = "Online"
	Card   PaymentMethod = "Card"
	Cash   PaymentMethod = "Cash"
)
