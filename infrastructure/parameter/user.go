package parameter

type SignUp struct {
	Email     string `json:"email"  form:"email" validate:"required,email"`
	Password  string `json:"password"  form:"password" validate:"required"`
	FirstName string `json:"first_name"  form:"first_name" validate:"required"`
	LastName  string `json:"last_name"  form:"last_name" validate:"required"`
}
type LoginIn struct {
	Email    string `json:"email"  form:"email" validate:"required,email"`
	Password string `json:"password"  form:"password" validate:"required"`
}
