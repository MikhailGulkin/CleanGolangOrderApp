package handlers

import (
	interfaces "github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/command"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductHandler struct {
	createProduct interfaces.CreateProduct
}

func (c *ProductHandler) CreateProduct(context *gin.Context) {
	var requestBody interfaces.CreateProductCommand
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
