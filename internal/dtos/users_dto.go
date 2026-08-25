package dtos

type GetUserByIdParams struct {
	ID int `validate:"required,gte=1"`
}

type CreateUserPayload struct {
	Username string `json:"username" validate:"required,min=6,max=100"`
	Email    string `json:"email" validate:"required,email,min=6,max=100"`
	Password string `json:"password" validate:"required,min=8,max=50"`
}

type GetUserByEmailPayload struct {
	Email string `json:"email" validate:"required,email,min=6,max=100"`
}

type GetUserByUsernameAndEmailPayload struct {
	Username string `json:"username" validate:"required,min=6,max=100"`
	Email    string `json:"email" validate:"required,email,min=6,max=100"`
}
