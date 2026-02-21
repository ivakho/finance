package update

import "context"

type usecaseUpdateTransaction interface {
	Update(ctx context.Context, id int, amount float64) error
}
