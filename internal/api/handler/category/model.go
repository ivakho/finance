package category

import "time"

type Category struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type CategoryTotal struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Total int64  `json:"total"`
}
