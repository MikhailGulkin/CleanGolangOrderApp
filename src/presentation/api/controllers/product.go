package controllers

import (
	interfaces "github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/command"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductController struct {
	Interactor interfaces.CreateProduct
}

func NewProductController(interactor interfaces.CreateProduct) ProductController {
	return ProductController{interactor}
}

func (c *ProductController) CreateProduct(context *gin.Context) {
	var requestBody interfaces.CreateProductCommand
	if err := context.BindJSON(&requestBody); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	err := c.Interactor.Create(requestBody)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	context.Status(http.StatusNoContent)
}
