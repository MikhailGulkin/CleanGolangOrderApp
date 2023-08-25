package commands

type CustomerCreateDTO struct {
	CustomerID string `json:"customerID" binding:"required"`
	EventID    string `json:"eventID" binding:"required"`
}
