package dtos

type GetUserByIdParams struct {
	ID int `validate:"required,gte=1"`
}
