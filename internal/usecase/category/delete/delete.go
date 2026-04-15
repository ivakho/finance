package delete

import (
	"context"
	"fmt"
)

func (u *Usecase) Delete(ctx context.Context, id int) error {
	if err := u.categoryRepo.DeleteCategory(ctx, id); err != nil {
		return fmt.Errorf("Failed to delete category: %w", err)
	}

	return nil
}
