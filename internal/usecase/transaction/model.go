package transaction

import (
	"time"
)

type Transaction struct {
	ID         int
	CategoryID int
	Amount     int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Transactions struct {
	Value []Transaction
	Total int64
}

type TransactionAdd struct {
	CategoryID int
	TxType     string
	Amount     int64
	CreatedAt  time.Time
}
