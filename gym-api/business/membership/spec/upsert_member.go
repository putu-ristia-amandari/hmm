package spec

import (
	"time"
)

type UpsertMemberCreateSpec struct {
	Name  string `validate:"required"`
	Image string
}

type UpsertMemberUpdateSpec struct {
	Name           string `validate:"required"`
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
