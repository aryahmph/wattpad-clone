package kaguya

type LoginRequest struct {
	Username string `validate:"required,min=5,max=100"`
	Password string `validate:"required,min=8,max=16"`
}
