package instructor

import "time"

type Instructor struct {
	ID          int
	Name        string
	Email       string
	Handphone   string
	Address     string
	City        string
	Province    string
	Nationality string
	Gender      string
	BirthOfDate time.Time
	Height      int
	Weight      int
	Photo       string
	Status      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
