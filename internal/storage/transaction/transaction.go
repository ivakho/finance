package transaction

import (
	"context"
	"database/sql"
	"errors"
	"finance/internal/model"
	postgresstorage "finance/internal/storage"
	"fmt"
	"time"
)

var ErrTransactionNotFound = errors.New("transaction not found")

type storage struct {
	postgresdb *postgresstorage.Postgres
}

func New(db *postgresstorage.Postgres) *storage {
	return &storage{postgresdb: db}
}

func (s *storage) AddTransaction(ctx context.Context, categoryID int, amount int64) error {
	query := "insert into transaction (category_id, amount, created_at, updated_at) values ($1, $2, $3, $4)"
	timeNow := time.Now()
	_, err := s.postgresdb.DB.ExecContext(ctx, query, categoryID, amount, timeNow, timeNow)
	if err != nil {
		return fmt.Errorf("ExecContext:%w", err)
	}

	return nil
}

func (s *storage) GetAllTransaction(ctx context.Context, categoryID int) ([]model.Transaction, error) {
	query := "select id, category_id, amount, created_at, updated_at from transaction where category_id = $1"

	rows, err := s.postgresdb.DB.QueryContext(ctx, query, categoryID)
	if err != nil {
		return nil, fmt.Errorf("QueryContext: %w", err)
	}
	defer rows.Close()

	var transactions []model.Transaction

	for rows.Next() {
		var transaction model.Transaction
		if err := rows.Scan(&transaction.ID, &transaction.CategoryID, &transaction.Amount, &transaction.CreatedAt, &transaction.UpdatedAt); err != nil {
			return nil, fmt.Errorf("Rows scan:%w", err)
		}
		transactions = append(transactions, transaction)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("Rows iteration:%w", rows.Err())
	}

	return transactions, nil
}

func (s *storage) GetTransaction(ctx context.Context, id int) (model.Transaction, error) {
	var transaction model.Transaction

	query := "select id, category_id, amount, created_at, updated_at from transaction where id = $1"

	err := s.postgresdb.DB.QueryRowContext(ctx, query, id).Scan(&transaction.ID, &transaction.CategoryID, &transaction.Amount, &transaction.CreatedAt, &transaction.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return transaction, ErrTransactionNotFound
		}
		return transaction, fmt.Errorf("QueryRowContext:%w", err)
	}

	return transaction, nil
}

func (s *storage) UpdateTransaction(ctx context.Context, id int, amount int64) error {
	query := "update transaction set amount = $1, updated_at = $2 where id = $3"

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
	query := "delete from transaction where id = $1"

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
