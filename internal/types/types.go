package types

type Student struct {
	ID    int    `json:"id"  validate:"-"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Class string `json:"class" validate:"required"`
}
