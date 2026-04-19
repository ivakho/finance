package category

import "time"

type Category struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt *time.Time
}

type CategoryTotal struct {
	ID    int
	Name  string
	Total int64
}
