package kaguya

type RegisterRequest struct {
	Username string `validate:"required,min=5,max=100"`
	Email    string `validate:"required,email,min=5,max=255"`
	Password string `validate:"required,min=8,max=16"`
}

type CheckSessionResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type CheckSessionRequest struct {
	Token string `validate:"required,max=3000"`
}
