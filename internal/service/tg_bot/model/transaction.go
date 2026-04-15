package model

import "time"

type Transaction struct {
	ID         int       `json:"id"`
	CategoryID int       `json:"category_id"`
	Type       string    `json:"type"`
	Amount     float64   `json:"amount"`
	CreatedAt  time.Time `json:"created_at"`
}
