package order

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/common/interfaces/persistence/filters"
	q "github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/common/interfaces/persistence/query"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/order/interfaces/command"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/order/interfaces/query"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/mediator"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/presentation/api/controllers/handlers"
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

	answer, err := c.mediator.Send(requestBody)
	if err != nil {
		context.Error(err)
		return
	}
	context.JSON(http.StatusNoContent, answer)
}
func (c *Handler) DeleteOrder(context *gin.Context) {
	var requestBody command.DeleteOrderCommand
	if err := context.BindJSON(&requestBody); err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	answer, err := c.mediator.Send(requestBody)
	if err != nil {
		context.Error(err)
		return
	}
	context.JSON(http.StatusNoContent, answer)
}
func (c *Handler) GetAllOrders(context *gin.Context) {
	Limit, Offset, Order := handlers.GetQueryParams(context)
	orders, err := c.mediator.Query(
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
	context.JSON(http.StatusOK, orders)
}
func (c *Handler) GetAllOrdersByUserID(context *gin.Context) {
	userID, errUUID := uuid.Parse(context.Param("userID"))
	if errUUID != nil {
		context.AbortWithError(http.StatusInternalServerError, errUUID)
		return
	}
	Limit, Offset, Order := handlers.GetQueryParams(context)
	orders, err := c.mediator.Query(
		query.GetAllOrderByUserIDQuery{
			UserID: userID,
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
	context.JSON(http.StatusOK, orders)
}
func (c *Handler) GetOrderByID(context *gin.Context) {
	orderID, errUUID := uuid.Parse(context.Param("id"))
	if errUUID != nil {
		context.AbortWithError(http.StatusInternalServerError, errUUID)
		return
	}
	orderByID, err := c.mediator.Query(
		query.GetOrderByIDQuery{ID: orderID},
	)
	if err != nil {
		context.Error(err)
		return
	}
	context.JSON(http.StatusOK, orderByID)
}
