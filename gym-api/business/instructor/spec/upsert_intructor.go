package spec

type UpsertInstructorCreateSpec struct {
	Name      string `validate:"required"`
	Email     string `validate:"required,email"`
	Gender    string `validate:"required"`
	City      string `validate:"required"`
	Handphone string `validate:"required"`
}

type UpsertInstructorUpdateSpec struct {
	Name        string `validate:"required"`
	Email       string `validate:"required,email"`
	Handphone   string `validate:"required"`
	Address     string
	City        string `validate:"required"`
	Province    string
	Nationality string
	Gender      string `validate:"required"`
	BirthOfDate string
	Height      int
	Weight      int
	Photo       string
	Status      bool
}
