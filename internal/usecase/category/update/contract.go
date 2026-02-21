package update

import (
	"context"
)

type categoryRepo interface {
	UpdateCategory(ctx context.Context, id int, name string) error
}
