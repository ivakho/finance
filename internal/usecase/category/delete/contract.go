package delete

import (
	"context"
)

type categoryRepo interface {
	DeleteCategory(ctx context.Context, id int) error
}
