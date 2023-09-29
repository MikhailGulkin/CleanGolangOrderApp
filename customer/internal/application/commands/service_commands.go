package commands

type CustomerCommands struct {
	CreateCustomer       *CreateCustomerHandler       `json:"createCustomer,omitempty"`
	UploadCustomerAvatar *UploadAvatarCustomerHandler `json:"uploadCustomerAvatar,omitempty"`
}

func NewCustomerCommands(
	createCustomer *CreateCustomerHandler,
	uploadCustomer *UploadAvatarCustomerHandler,
) *CustomerCommands {
	return &CustomerCommands{
		CreateCustomer:       createCustomer,
		UploadCustomerAvatar: uploadCustomer,
	}
}
