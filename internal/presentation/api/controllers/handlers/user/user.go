package user

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/user/interfaces/command"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	createUser command.CreateUser
}

func (c *Handler) CreateUser(context *gin.Context) {
	var requestBody command.CreateUserCommand
	if err := context.BindJSON(&requestBody); err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err := c.createUser.Create(requestBody)
	if err != nil {
		context.Error(err)
		return
	}
	context.Status(http.StatusNoContent)
}
