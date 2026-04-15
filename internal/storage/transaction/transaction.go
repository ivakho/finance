package transaction

import (
	"context"
	"errors"
	postgresstorage "finance/internal/storage"
	"fmt"
	"time"
)

var ErrTransactionNotFound = errors.New("transaction not found")

const defaultLimit = 50

type storage struct {
	postgresdb *postgresstorage.Postgres
}

func New(db *postgresstorage.Postgres) *storage {
	return &storage{postgresdb: db}
}

func (s *storage) AddTransaction(ctx context.Context, categoryID int, amount int64, createdAt time.Time) error {
	query := "insert into transactions (category_id, amount, created_at, updated_at) values ($1, $2, $3, $4)"
	_, err := s.postgresdb.DB.ExecContext(ctx, query, categoryID, amount, createdAt, nil)
	if err != nil {
		return fmt.Errorf("ExecContext:%w", err)
	}

	return nil
}

func (s *storage) GetTransactionByID(ctx context.Context, id int) (TransactionByID, error) {
	query := `
		select 
			t.id,
			t.category_id,
			c.name,
			t.amount,
			t.created_at,
			t.updated_at
		from transactions t
		join category c on t.category_id = c.id
		where t.id = $1
	`
	var transaction TransactionByID

	err := s.postgresdb.DB.QueryRowContext(ctx, query, id).
		Scan(&transaction.ID, &transaction.CategoryID, &transaction.CategoryName, &transaction.Amount, &transaction.CreatedAt, &transaction.UpdatedAt)

	if err != nil {
		return TransactionByID{}, fmt.Errorf("QueryRowContext: %w", err)
	}

	return transaction, nil
}

func (s *storage) GetTransaction(ctx context.Context, categoryID int, dateFrom, dateTo time.Time) ([]Transaction, error) {

	query := "select id, category_id, amount, created_at, updated_at from transactions where category_id=$1 and created_at::date >= $2 and created_at::date <= $3"

	rows, err := s.postgresdb.DB.QueryContext(ctx, query, categoryID, dateFrom, dateTo)
	if err != nil {
		return nil, fmt.Errorf("QueryContext:%w", err)
	}
	defer rows.Close()

	var transactions []Transaction
	for rows.Next() {
		var transaction Transaction

		if err := rows.Scan(&transaction.ID, &transaction.CategoryID, &transaction.Amount, &transaction.CreatedAt, &transaction.UpdatedAt); err != nil {
			return nil, fmt.Errorf("RowsScan:%w", err)
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (s *storage) UpdateTransaction(ctx context.Context, id int, amount int64) error {
	query := "update transactions set amount = $1, updated_at = $2 where id = $3"

	result, err := s.postgresdb.DB.ExecContext(ctx, query, amount, time.Now(), id)
	if err != nil {
		return fmt.Errorf("ExecContext:%w", err)
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("RowsAffected:%w", err)
	}

	if rowAffected == 0 {
		return ErrTransactionNotFound
	}

	return nil
}

func (s *storage) DeleteTransaction(ctx context.Context, id int) error {
	query := "delete from transactions where id = $1"

	result, err := s.postgresdb.DB.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("ExecContext:%w", err)
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("RowsAffected:%w", err)
	}

	if rowAffected == 0 {
		return ErrTransactionNotFound
	}

	return nil
}
