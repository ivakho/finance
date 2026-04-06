package delete

import "context"

type usecaseDeleteCategory interface {
	Delete(ctx context.Context, id int) error
}
