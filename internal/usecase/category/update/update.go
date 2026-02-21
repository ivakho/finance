package update

import (
	"context"
	"fmt"
)

func (u *Usecase) Update(ctx context.Context, id int, name string) error {
	if err := u.categoryRepo.UpdateCategory(ctx, id, name); err != nil {
		return fmt.Errorf("Failed to edit category: %w", err)
	}

	return nil
}
