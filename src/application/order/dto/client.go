package dto

import "github.com/google/uuid"

type Client struct {
	ClientID   uuid.UUID
	ClientName string
}
