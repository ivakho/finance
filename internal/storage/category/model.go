package category

import "time"

type Category struct {
	ID        int        `db:"id"`
	Name      string     `db:"name"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
}

type CategoryTotal struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Total int64  `db:"total"`
}
