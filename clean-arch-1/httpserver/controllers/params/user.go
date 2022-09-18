package params

type UserCreateRequest struct {
	Email string `validate:"required,email"`
	Password string `validate:"required"`
}
