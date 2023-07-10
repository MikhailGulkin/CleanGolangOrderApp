package handlers

import (
	f "github.com/MikhailGulkin/simpleGoOrderApp/src/application/common/interfaces/persistence/filters"
	q "github.com/MikhailGulkin/simpleGoOrderApp/src/application/common/interfaces/persistence/query"
	command2 "github.com/MikhailGulkin/simpleGoOrderApp/src/application/product/interfaces/command"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/product/interfaces/query"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	createProduct     command2.CreateProduct
	getAllProducts    query.GetAllProducts
	getProductByName  query.GetProductByName
	updateProductName command2.UpdateProductName
}

func (c *ProductHandler) CreateProduct(context *gin.Context) {
	var requestBody command2.CreateProductCommand
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
func (c *ProductHandler) UpdateProductName(context *gin.Context) {
	productID := context.Param("productID")

	var requestBody command2.UpdateProductNameCommand
	if err := context.BindJSON(&requestBody); err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	parse, err := uuid.Parse(productID)
	if err != nil {
		context.Error(err)
	}
	requestBody.ProductID = parse
	err = c.updateProductName.Update(requestBody)
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
		query.GetAllProductsQuery{
			BaseListQueryParams: q.BaseListQueryParams{
				Limit:  uint(Limit),
				Offset: uint(Offset),
				Order:  f.ConvertToOrder(Order)},
		},
	)
	if err != nil {
		context.Error(err)
		return
	}
	context.JSON(http.StatusOK, products)
}
func (c *ProductHandler) GetProductByName(context *gin.Context) {
	productName := context.Param("productName")
	productByName, err := c.getProductByName.Get(
		query.GetProductByNameQuery{Name: productName},
	)
	if err != nil {
		context.Error(err)
		return
	}
	context.JSON(http.StatusOK, productByName)
}
