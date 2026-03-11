package model

import (
	"time"
)

type Transaction struct {
	ID         int       `json:"id"`
	CategoryID int       `json:"category_id"`
	Amount     int64       `json:"amount"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
