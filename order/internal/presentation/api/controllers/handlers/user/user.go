package user

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/user/interfaces/command"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/mediator"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	mediator mediator.Mediator
}

func (c *Handler) CreateUser(context *gin.Context) {
	var requestBody command.CreateUserCommand
	if err := context.BindJSON(&requestBody); err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err := c.mediator.Send(requestBody)
	if err != nil {
		context.Error(err)
		return
	}
	context.Status(http.StatusNoContent)
}
