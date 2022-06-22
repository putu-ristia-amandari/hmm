package newsletter

import (
	"gym-membership/business/category"
	"time"
)

type NewsLetter struct {
	ID         int
	CategoryID int
	Title      string
	Image      string
	Body       string
	CreatedBy  string
	UpdatedBy  string
	CreatedAt  time.Time
	UpdatedAt  time.Time

	Category category.Category
}
