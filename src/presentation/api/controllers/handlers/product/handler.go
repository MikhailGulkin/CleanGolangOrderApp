package product

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/common/interfaces/persistence/filters"
	q "github.com/MikhailGulkin/simpleGoOrderApp/src/application/common/interfaces/persistence/query"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/product/interfaces/command"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/product/interfaces/query"
	commandbus "github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/commandBus"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/controllers/handlers"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type Handler struct {
	bus               commandbus.CommandBus
	getAllProducts    query.GetAllProducts
	getProductByName  query.GetProductByName
	updateProductName command.UpdateProductName
}

func (c *Handler) CreateProduct(context *gin.Context) {
	var requestBody command.CreateProductCommand
	if err := context.BindJSON(&requestBody); err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err := c.bus.Execute(requestBody)
	if err != nil {
		context.Error(err)
		return
	}
	context.Status(http.StatusNoContent)
}
func (c *Handler) UpdateProductName(context *gin.Context) {
	productID := context.Param("productID")

	var requestBody command.UpdateProductNameCommand
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
func (c *Handler) GetAllProducts(context *gin.Context) {
	Limit, Offset, Order := handlers.GetQueryParams(context)
	products, err := c.getAllProducts.Get(
		query.GetAllProductsQuery{
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
func (c *Handler) GetProductByName(context *gin.Context) {
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
