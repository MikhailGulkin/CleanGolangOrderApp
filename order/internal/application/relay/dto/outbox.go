package dto

import "github.com/google/uuid"

type Message struct {
	ID       uuid.UUID
	Exchange string
	Route    string
	Payload  []byte
}
