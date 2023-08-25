package handlers

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type (
	Limit  int
	Offset int
	Order  string
)

func GetQueryParams(context *gin.Context) (Limit, Offset, Order) {
	limit, _ := strconv.Atoi(context.DefaultQuery("limit", "1000"))
	offset, _ := strconv.Atoi(context.DefaultQuery("offset", "0"))
	order := context.DefaultQuery("order", "asc")
	return Limit(limit), Offset(offset), Order(order)
}
