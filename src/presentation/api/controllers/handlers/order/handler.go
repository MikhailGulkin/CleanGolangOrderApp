package order

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/common/interfaces/persistence/filters"
	q "github.com/MikhailGulkin/simpleGoOrderApp/src/application/common/interfaces/persistence/query"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/command"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/query"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strconv"
)

type Handler struct {
	createOrder  command.CreateOrder
	getAllOrders query.GetAllOrders
	getOrderByID query.GetOrderByID
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
func (c *Handler) GetAllOrders(context *gin.Context) {
	Limit, _ := strconv.Atoi(context.DefaultQuery("limit", "1000"))
	Offset, _ := strconv.Atoi(context.DefaultQuery("offset", "0"))
	Order := context.DefaultQuery("order", "asc")

	products, err := c.getAllOrders.Get(
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
	productByName, err := c.getOrderByID.Get(
		query.GetOrderByIDQuery{ID: orderID},
	)
	if err != nil {
		context.Error(err)
		return
	}
	context.JSON(http.StatusOK, productByName)
}
