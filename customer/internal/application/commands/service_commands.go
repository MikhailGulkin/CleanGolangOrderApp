package commands

type CustomerCommands struct {
	CreateCustomer CreateCustomerHandler `json:"createCustomer" binding:"required"`
}

func NewCustomerCommands(createCustomer CreateCustomerHandler) *CustomerCommands {
	return &CustomerCommands{
		CreateCustomer: createCustomer,
	}
}
