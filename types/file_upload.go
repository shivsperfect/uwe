package types

import (
	"time"

	"github.com/google/uuid"
)

type FileType int

const (
	FileTypeSubscriptions FileType = iota
)

type FileUpload struct {
	ID         uuid.UUID      `json:"id"`
	CustomerID uuid.UUID      `json:"customer_id"`
	Type       FileType       `json:"type"`
	Mapping    map[string]int `json:"mapping" bun:"default:{}"`
	CreatedAt  time.Time      `json:"created_at"`
}
