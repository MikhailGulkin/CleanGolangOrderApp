package command

import "github.com/google/uuid"

type CreateUserCommand struct {
	UserID    uuid.UUID `json:"userID" binding:"required"`
	Username  string    `json:"username" binding:"required"`
	AddressID uuid.UUID `json:"addressID" binding:"required"`
}
type CreateUser interface {
	Create(command CreateUserCommand) error
}
