package category

import (
	"context"
	"errors"
	modelcategory "finance/internal/model"
	postgresstorage "finance/internal/storage"
	"fmt"
	"time"
)

var ErrCategoryNotFound = errors.New("category not found")

type storage struct {
	postgresdb *postgresstorage.Postgres
}

func New(db *postgresstorage.Postgres) *storage {
	return &storage{
		postgresdb: db,
	}
}

func (s *storage) AddCategory(ctx context.Context, name string) error {
	query := "insert into category (name, created_at) values ($1, $2)"

	// Говорят, что лучше использовать ExecContext, т.к. нам не надо возвращать row
	// еще пишут, что без Scan() соединение зависнет и не вернётся в пул
	row := s.postgresdb.DB.QueryRowContext(ctx, query, name, time.Now())
	if row.Err() != nil {
		return fmt.Errorf("QueryRowContext: %w", row.Err())
	}
	return nil
}

func (s *storage) GetAllCategory(ctx context.Context) ([]modelcategory.Category, error) {
	query := "select id, name, created_at, updated_at from category"
	rows, err := s.postgresdb.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("QueryContext: %w", err)
	}
	defer rows.Close()
	categories := make([]modelcategory.Category, 0)

	for rows.Next() {
		var category modelcategory.Category
		if err := rows.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.UpdatedAt); err != nil {
			return nil, fmt.Errorf("Rows scan: %w", err)
		}
		categories = append(categories, category)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("Rows iteration: %w", rows.Err())
	}

	return categories, nil
}

func (s *storage) UpdateCategory(ctx context.Context, id int, name string) error {
	query := "update category set name = $1, updated_at = $2 where id = $3"

	result, err := s.postgresdb.DB.ExecContext(ctx, query, name, time.Now(), id)
	if err != nil {
		return fmt.Errorf("ExecContext:%w", err)
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("RowsAffected:%w", err)
	}

	if rowAffected == 0 {
		return ErrCategoryNotFound
	}

	return nil
}

func (s *storage) DeleteCategory(ctx context.Context, id int) error {
	query := "delete from category where id = $1"

	result, err := s.postgresdb.DB.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("ExecContext:%w", err)
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("RowsAffected:%w", err)
	}

	if rowAffected == 0 {
		return ErrCategoryNotFound
	}

	return nil
}
