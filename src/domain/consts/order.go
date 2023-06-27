package consts

type OrderStatus int

const (
	New OrderStatus = iota
	InProcessing
	Processed
	Delivered
	Canceled
)

type PaymentMethod int

const (
	Online PaymentMethod = iota
	Card
	Cash
)
