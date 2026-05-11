package types

type Student struct {
	ID    int64  `json:"id"  validate:"-"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Class string `json:"class" validate:"required"`
}
