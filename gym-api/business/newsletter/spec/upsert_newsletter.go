package spec

type UpsertNewsLetterCreateSpec struct {
	Title     string `validate:"required"`
	Image     string `validate:"required"`
	Body      string `validate:"required"`
	CreatedBy string `validate:"required"`
	UpdatedBy string `validate:"required"`
}

type UpsertNewsLetterUpdateSpec struct {
	Title     string `validate:"required"`
	Image     string `validate:"required"`
	Body      string `validate:"required"`
	CreatedBy string `validate:"required"`
	UpdatedBy string `validate:"required"`
}
