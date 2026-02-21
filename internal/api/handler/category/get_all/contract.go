package get_all

import (
	"context"
	"finance/internal/model"
)

type usecaseGetAllCategory interface {
	GetAll(ctx context.Context) ([]model.Category, error)
}
