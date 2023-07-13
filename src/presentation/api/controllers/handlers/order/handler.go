package order

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/command"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	createOrder command.CreateOrder
}

func (c *Handler) CreateOrder(context *gin.Context) {
	var requestBody command.CreateOrderCommand
	if err := context.BindJSON(&requestBody); err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err := c.createOrder.Create(requestBody)
	if err != nil {
		context.Error(err)
		return
	}
	context.Status(http.StatusNoContent)
}
