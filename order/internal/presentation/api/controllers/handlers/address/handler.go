package address

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/address/interfaces/command"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/mediator"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	mediator mediator.Mediator
}

func (c *Handler) CreateAddress(context *gin.Context) {
	var requestBody command.CreateAddressCommand
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
