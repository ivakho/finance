package update

import "context"

type usecaseUpdateCategory interface {
	Update(ctx context.Context, id int, name string) error
}
