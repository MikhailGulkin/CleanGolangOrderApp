package dto

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/common/dto"
	"github.com/google/uuid"
	"time"
)

type Product struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Price        float64   `json:"price"`
	Discount     int32     `json:"discount"`
	Quantity     int32     `json:"quantity"`
	Description  string    `json:"description"`
	Availability bool      `json:"availability"`
	CreatedAt    time.Time `json:"created_at"`
}

type Products struct {
	Products []Product `json:"products,omitempty"`
	dto.BaseSequence
}
