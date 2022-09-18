package services

import (
	"clean-arch-1/httpserver/controllers/params"
	"clean-arch-1/httpserver/controllers/views"
	"clean-arch-1/httpserver/repositories/models"

	"golang.org/x/crypto/bcrypt"
)

var users = []models.User{
	{
		ID:       1,
		Email: "jovinlidan@gmail.com",
		Password: "jovinlidan",
	},
	{
		ID:       2,
		Email: "jovinlidann@gmail.com",
		Password: "jovinlidann",
	},
	{
		ID:       3,
		Email: "jovinlidannn@gmail.com",
		Password: "jovinlidannn",
	},
}

func CreateUser(req *params.UserCreateRequest) *views.Response {
	var user models.User

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 4)
	if err != nil {
		return views.BadRequestError(err)
	}

	user.ID = len(users)+1
	user.Email = req.Email
	user.Password = string(hash)


	users = append(users, user)

	res := views.SuccessCreateResponse(user, "Success Create User !!!")

	return res
}

func GetUsers() *views.Response {
	res := views.SuccessCreateResponse(users, "Success Get Users !!!")
	return res
}
