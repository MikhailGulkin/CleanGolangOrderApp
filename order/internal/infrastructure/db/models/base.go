package models

import (
	"github.com/google/uuid"
	"time"
)

type Base struct {
	ID        uuid.UUID  `gorm:"index;type:uuid;primary_key;default:uuid_generate_v4()"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP; not null"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP; not null"`
	DeletedAt *time.Time `gorm:"default:null"`
}
