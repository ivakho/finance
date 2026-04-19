package transaction

import "time"

type Transaction struct {
	ID        int       `json:"id"`
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type TransactionByID struct {
	ID           int        `json:"id"`
	CategoryID   int        `json:"category_id"`
	CategoryName string     `json:"category_name"`
	Amount       int64      `json:"amount"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
}

type Transactions []Transaction
