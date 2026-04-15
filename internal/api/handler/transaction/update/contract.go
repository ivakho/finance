package update

import "context"

type usecaseUpdateTransaction interface {
	Update(ctx context.Context, id int, amount int64) error
}
