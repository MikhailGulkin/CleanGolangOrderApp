package product

import (
	"encoding/json"
	"fmt"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/entities/product"
	"github.com/MikhailGulkin/simpleGoOrderApp/tests/integration/product/utils"
	"net/http"
)

func (suite *TestSuite) TestGetProductByName() {
	utils.CreateProductInDB(&suite.ProductModel, suite.Server.DB)

	resp, err := http.Get(fmt.Sprintf("%s/products/%s", suite.Server.URL, suite.ProductModel.Name)) //nolint:bodyclose
	if err != nil {
		suite.Fail(fmt.Sprintf("Expected no errorHandler, got %v", err))
	}
	suite.Equal(http.StatusOK, resp.StatusCode)

	decoder := json.NewDecoder(resp.Body)
	var data product.Product
	_ = decoder.Decode(&data)

	suite.Equal(data.Name, suite.ProductModel.Name)
}
func (suite *TestSuite) TestNotExistProductName() {
	resp, err := http.Get(fmt.Sprintf("%s/products/%s", suite.Server.URL, suite.ProductModel.Name)) //nolint:bodyclose
	if err != nil {
		suite.Fail(fmt.Sprintf("Expected no errorHandler, got %v", err))
	}
	suite.Equal(http.StatusNotFound, resp.StatusCode)
}
