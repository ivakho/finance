package transaction

import "time"

type Transaction struct {
	ID        int       `json:"id"`
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type Transactions []Transaction
