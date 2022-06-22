package membership

import "time"

type Membership struct {
	ID             int
	Name           string
	Image          string
	LimitedClass   int
	LimitedTime    int
	LimitedContent int
	Description    string
	Details        string
	Price          float64
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
