package handlers

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/command"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/persistence/filters"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/query"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/query/product"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	createProduct  command.CreateProduct
	getAllProducts product.GetAllProducts
}

func (c *ProductHandler) CreateProduct(context *gin.Context) {
	var requestBody command.CreateProductCommand
	if err := context.BindJSON(&requestBody); err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err := c.createProduct.Create(requestBody)
	if err != nil {
		context.Error(err)
		return
	}
	context.Status(http.StatusNoContent)
}
func (c *ProductHandler) GetALlProducts(context *gin.Context) {
	Limit, _ := strconv.Atoi(context.DefaultQuery("limit", "1000"))
	Offset, _ := strconv.Atoi(context.DefaultQuery("offset", "0"))
	Order := context.DefaultQuery("order", "asc")

	products, err := c.getAllProducts.Get(
		product.GetAllProductsQuery{
			BaseListQueryParams: query.BaseListQueryParams{
				Limit:  uint(Limit),
				Offset: uint(Offset),
				Order:  filters.ConvertToOrder(Order)},
		},
	)
	if err != nil {
		context.Error(err)
		return
	}
	context.JSON(http.StatusOK, products)
}
