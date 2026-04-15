package delete

import "context"

type usecaseDeleteTransaction interface {
	Delete(ctx context.Context, id int) error
}
