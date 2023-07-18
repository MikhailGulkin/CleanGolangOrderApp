package order

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/common/interfaces/persistence/filters"
	q "github.com/MikhailGulkin/simpleGoOrderApp/internal/application/common/interfaces/persistence/query"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/interfaces/command"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/interfaces/query"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/mediator"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/presentation/api/controllers/handlers"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type Handler struct {
	mediator mediator.Mediator
}

func (c *Handler) CreateOrder(context *gin.Context) {
	var requestBody command.CreateOrderCommand
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
func (c *Handler) GetAllOrders(context *gin.Context) {
	Limit, Offset, Order := handlers.GetQueryParams(context)
	products, err := c.mediator.Query(
		query.GetAllOrderQuery{
			BaseListQueryParams: q.BaseListQueryParams{
				Limit:  uint(Limit),
				Offset: uint(Offset),
				Order:  filters.BaseOrder(Order),
			},
		},
	)
	if err != nil {
		context.Error(err)
		return
	}
	context.JSON(http.StatusOK, products)
}
func (c *Handler) GetOrderByID(context *gin.Context) {
	orderID, errUUID := uuid.Parse(context.Param("id"))
	if errUUID != nil {
		context.AbortWithError(http.StatusInternalServerError, errUUID)
		return
	}
	productByName, err := c.mediator.Query(
		query.GetOrderByIDQuery{ID: orderID},
	)
	if err != nil {
		context.Error(err)
		return
	}
	context.JSON(http.StatusOK, productByName)
}
