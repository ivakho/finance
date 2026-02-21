package get

import (
	"context"
	"finance/internal/model"
)

type usecaseGetTransaction interface {
	Get(ctx context.Context, id int) (model.Transaction, error)
}
