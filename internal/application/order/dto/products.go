package dto

import "github.com/google/uuid"

type Product struct {
	ProductID   uuid.UUID `json:"productID"`
	Name        string    `json:"name"`
	ActualPrice float64   `json:"actualPrice"`
}
