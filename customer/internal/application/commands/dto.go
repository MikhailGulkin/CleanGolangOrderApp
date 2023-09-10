package commands

type CustomerCreateDTO struct {
	CustomerID string `json:"customerID" binding:"required"`
	EventID    string `json:"eventID" binding:"required"`
}

type UploadAvatarDTO struct {
	AvatarUri  string `json:"avatarUri" binding:"required"`
	CustomerID string `json:"customerID" binding:"required"`
	EventID    string `json:"eventID" binding:"required"`
}
