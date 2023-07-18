package user

import c "github.com/MikhailGulkin/simpleGoOrderApp/internal/application/user/interfaces/command"

func NewUserHandler(
	createUser c.CreateUser,
) Handler {
	return Handler{
		createUser: createUser,
	}
}
