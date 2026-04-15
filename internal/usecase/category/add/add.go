package add

import (
	"context"
	"fmt"
)

func (u *Usecase) Add(ctx context.Context, name string) error {
	if err := u.categoryRepo.AddCategory(ctx, name); err != nil {
		return fmt.Errorf("Failed to add category: %w", err)
	}

	return nil
}
