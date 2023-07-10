package product

import (
	"fmt"
	"github.com/MikhailGulkin/simpleGoOrderApp/tests/integration/product/utils"
	"net/http"
	"strings"
)

var ZeroInt64 = int64(0)

func (suite *TestSuite) TestSuccessProductCreate() {
	resp, err := http.Post(fmt.Sprintf("%s/products", suite.Server.URL), "application/json", utils.CreateValidByteProduct()) //nolint:bodyclose

	if err != nil {
		suite.Fail(fmt.Sprintf("Expected no errorHandler, got %v", err))
	}
	suite.Equal(http.StatusNoContent, resp.StatusCode)
	suite.Equal(resp.ContentLength, ZeroInt64)
}
func (suite *TestSuite) TestInvalidDiscount() {
	resp, err := http.Post(fmt.Sprintf("%s/products", suite.Server.URL), "application/json", utils.CreateInvalidDiscountByteProduct()) //nolint:bodyclose

	if err != nil {
		suite.Fail(fmt.Sprintf("Expected no errorHandler, got %v", err))
	}
	suite.Equal(http.StatusBadRequest, resp.StatusCode)
	suite.Equal(resp.ContentLength, ZeroInt64)
}
func (suite *TestSuite) TestInvalidPrice() {
	resp, err := http.Post(fmt.Sprintf("%s/products", suite.Server.URL), "application/json", utils.CreateInvalidPriceByteProduct()) //nolint:bodyclose

	if err != nil {
		suite.Fail(fmt.Sprintf("Expected no errorHandler, got %v", err))
	}
	suite.Equal(http.StatusBadRequest, resp.StatusCode)
	suite.Equal(resp.ContentLength, ZeroInt64)
}
func (suite *TestSuite) TestNotExistProductWhenUpdate() {
	req, err := http.NewRequest(
		http.MethodPut,
		fmt.Sprintf("%s/products/%s/productName", suite.Server.URL, suite.ProductModel.ID),
		strings.NewReader(fmt.Sprintf("{\"productName\": \"%s\"}", "testNameUpdateNotExist")),
	)
	if err != nil {
		suite.Fail(fmt.Sprintf("Expected no errorHandler, got %v", err))
	}

	client := &http.Client{}
	resp, errClient := client.Do(req) //nolint:bodyclose
	if errClient != nil {
		suite.Fail(fmt.Sprintf("Expected no errorHandler, got %v", errClient))
	}
	suite.Equal(http.StatusNotFound, resp.StatusCode)
	suite.Equal(resp.ContentLength, ZeroInt64)
}
func (suite *TestSuite) TestSuccessUpdateProductName() {
	utils.CreateProductInDB(&suite.ProductModel, suite.Server.DB)
	req, err := http.NewRequest(
		http.MethodPut,
		fmt.Sprintf("%s/products/%s/productName", suite.Server.URL, suite.ProductModel.ID),
		strings.NewReader(fmt.Sprintf("{\"productName\": \"%s\"}", "TestNameUpdateSuccess")),
	)
	if err != nil {
		suite.Fail(fmt.Sprintf("Expected no errorHandler, got %v", err))
	}

	client := &http.Client{}
	resp, errClient := client.Do(req) //nolint:bodyclose
	if errClient != nil {
		suite.Fail(fmt.Sprintf("Expected no errorHandler, got %v", errClient))
	}
	suite.Equal(http.StatusNoContent, resp.StatusCode)
	suite.Equal(resp.ContentLength, ZeroInt64)
}
func (suite *TestSuite) TestIncorrectUpdateProductName() {
	utils.CreateProductInDB(&suite.ProductModel, suite.Server.DB)
	req, err := http.NewRequest(
		http.MethodPut,
		fmt.Sprintf("%s/products/%s/productName", suite.Server.URL, suite.ProductModel.ID),
		strings.NewReader(fmt.Sprintf("{\"productName\": \"%s\"}", "testNameUpdateSuccess")),
	)
	if err != nil {
		suite.Fail(fmt.Sprintf("Expected no errorHandler, got %v", err))
	}

	client := &http.Client{}
	resp, errClient := client.Do(req) //nolint:bodyclose
	if errClient != nil {
		suite.Fail(fmt.Sprintf("Expected no errorHandler, got %v", errClient))
	}
	suite.Equal(http.StatusBadRequest, resp.StatusCode)
	suite.Equal(resp.ContentLength, ZeroInt64)
}
