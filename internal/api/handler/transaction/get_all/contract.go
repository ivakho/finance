package get_all

import (
	"context"
	"finance/internal/model"
)

type usecaseGetAllTransaction interface {
	GetAll(ctx context.Context, categoryID int) ([]model.Transaction, error)
}
