package add

import (
	"context"
)

type categoryRepo interface {
	AddCategory(ctx context.Context, name string) error
}
