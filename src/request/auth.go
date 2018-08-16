package request

type Signin struct {
	Name     string `form:"name" validate:"omitempty,max=100"`
	Password string `form:"password" validate:"required,min=6,max=100"`
	Email    string `form:"email" validate:"omitempty,email,max=100"`
}

type Signup struct {
	Name     string `form:"name" validate:"required,max=100"`
	Password string `form:"password" validate:"required,min=6,max=100"`
	Email    string `form:"email" validate:"required,email,max=100"`
}
