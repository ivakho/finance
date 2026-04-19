package transaction

import "time"

type Transaction struct {
	ID         int        `db:"id"`
	CategoryID int        `db:"category_id"`
	Amount     int64      `db:"amount"`
	CreatedAt  time.Time  `db:"created_at"`
	UpdatedAt  *time.Time `db:"updated_at"`
}

type TransactionByID struct {
	ID           int
	CategoryID   int
	CategoryName string
	Amount       int64
	CreatedAt    time.Time
	UpdatedAt    *time.Time
}
