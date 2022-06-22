package category

import "time"

type Category struct {
	ID          int
	Name        string
	Image       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
