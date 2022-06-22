package spec

type UpsertContentCreateSpec struct {
	Title       string `validate:"required"`
	LinkVideo   string `validate:"required"`
	Description string `validate:"required"`
	CreatedBy   string `validate:"required"`
}

type UpsertContentUpdateSpec struct {
	Title       string `validate:"required"`
	Image       string
	LinkVideo   string `validate:"required"`
	Description string `validate:"required"`
	CreatedBy   string `validate:"required"`
	UpdatedBy   string `validate:"required"`
}
