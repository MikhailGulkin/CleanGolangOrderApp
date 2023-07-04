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
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	err := c.createProduct.Create(requestBody)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	context.Status(http.StatusNoContent)
}
