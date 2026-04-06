package add

import "context"

type usecaseCategoryAdd interface {
	Add(ctx context.Context, name string) error
}
