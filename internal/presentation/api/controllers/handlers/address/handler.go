package address

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/address/interfaces/command"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	createAddress command.CreateAddress
}

func (c *Handler) CreateAddress(context *gin.Context) {
	var requestBody command.CreateAddressCommand
	if err := context.BindJSON(&requestBody); err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err := c.createAddress.Create(requestBody)
	if err != nil {
		context.Error(err)
		return
	}
	context.Status(http.StatusNoContent)
}
