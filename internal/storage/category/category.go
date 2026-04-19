package category

import (
	"context"
	"errors"
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

	_, err := s.postgresdb.DB.ExecContext(ctx, query, name, time.Now())
	if err != nil {
		return fmt.Errorf("QueryRowContext: %w", err)
	}

	return nil
}

func (s *storage) GetCategoryByID(ctx context.Context, id int) (Category, error) {
	query := "select id, name, created_at, updated_at from category where id = $1"

	var category Category

	err := s.postgresdb.DB.QueryRowContext(ctx, query, id).
		Scan(&category.ID, &category.Name, &category.CreatedAt, &category.UpdatedAt)

	if err != nil {
		return Category{}, fmt.Errorf("QueryRowContext: %w", err)
	}

	return category, nil
}

func (s *storage) GetAllCategory(ctx context.Context) ([]Category, error) {
	query := "select id, name, created_at, updated_at from category"
	rows, err := s.postgresdb.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("QueryContext: %w", err)
	}
	defer rows.Close()
	categories := make([]Category, 0)

	for rows.Next() {
		var category Category
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

func (s *storage) GetCategoryIncomeTotal(ctx context.Context, dateFrom, dateTo time.Time) ([]CategoryTotal, error) {
	query := `
        select 
            c.id as category_id,
            c.name as category_name,
            coalesce(sum(t.amount), 0) as total
        from category c
        left join transactions t 
            on c.id = t.category_id 
            and t.created_at::date BETWEEN $1 and $2
						and t.amount > 0
						group by c.id, c.name order by c.id
    `

	rows, err := s.postgresdb.DB.QueryContext(ctx, query, dateFrom, dateTo)
	if err != nil {
		return nil, fmt.Errorf("QueryContext: %w", err)
	}
	defer rows.Close()

	var result []CategoryTotal
	for rows.Next() {
		var ct CategoryTotal
		if err := rows.Scan(&ct.ID, &ct.Name, &ct.Total); err != nil {
			return nil, fmt.Errorf("Scan: %w", err)
		}
		result = append(result, ct)
	}

	return result, nil
}

func (s *storage) GetCategoryExpenseTotal(ctx context.Context, dateFrom, dateTo time.Time) ([]CategoryTotal, error) {
	query := `
        select 
            c.id as category_id,
            c.name as category_name,
            coalesce(sum(t.amount), 0) as total
        from category c
        left join transactions t 
            on c.id = t.category_id 
            and t.created_at::date BETWEEN $1 and $2
						and t.amount < 0
						group by c.id, c.name order by c.id
    `

	rows, err := s.postgresdb.DB.QueryContext(ctx, query, dateFrom, dateTo)
	if err != nil {
		return nil, fmt.Errorf("QueryContext: %w", err)
	}
	defer rows.Close()

	var result []CategoryTotal
	for rows.Next() {
		var ct CategoryTotal
		if err := rows.Scan(&ct.ID, &ct.Name, &ct.Total); err != nil {
			return nil, fmt.Errorf("Scan: %w", err)
		}
		result = append(result, ct)
	}

	return result, nil
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
	tx, err := s.postgresdb.DB.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("BeginTx: %w", err)
	}
	defer tx.Rollback()

	deleteTransactionsQuery := "DELETE FROM transactions WHERE category_id = $1"
	_, err = tx.ExecContext(ctx, deleteTransactionsQuery, id)
	if err != nil {
		return fmt.Errorf("Delete transactions: %w", err)
	}

	deleteCategoryQuery := "DELETE FROM category WHERE id = $1"
	result, err := tx.ExecContext(ctx, deleteCategoryQuery, id)
	if err != nil {
		return fmt.Errorf("Delete category: %w", err)
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("RowsAffected: %w", err)
	}

	if rowAffected == 0 {
		return ErrCategoryNotFound
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("Commit: %w", err)
	}

	return nil
}
