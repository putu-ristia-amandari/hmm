package content

import (
	"gym-membership/business/category"
	"time"
)

type Content struct {
	ID          int
	CategoryID  int
	Title       string
	Image       string
	LinkVideo   string
	Description string
	CreatedBy   string
	UpdatedBy   string
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Category category.Category
}
