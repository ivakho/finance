package transaction

import "time"

type Transaction struct {
	ID         int        `db:"id"`
	CategoryID int        `db:"category_id"`
	Type       string     `db:"type"`
	Amount     int64      `db:"amount"`
	CreatedAt  time.Time  `db:"created_at"`
	UpdatedAt  *time.Time `db:"updated_at"`
}

type TransactionFilter struct {
	CategoryID int
	Type       string
	DateFrom   *time.Time
	DateTo     *time.Time
	Limit      int
}
