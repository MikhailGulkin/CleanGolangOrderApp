package api

import (
	"encoding/json"
	"fmt"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/dto"
	"github.com/MikhailGulkin/simpleGoOrderApp/tests/integration/api/order/utils"
	"net/http"
)

func (suite *TestSuite) TestGetAllOrders() {
	utils.CreateOrderInDB(&suite.OrderModel, suite.Server.DB)
	resp, err := http.Get(fmt.Sprintf("%s/orders/", suite.Server.URL)) //nolint:bodyclose
	if err != nil {
		suite.Fail(fmt.Sprintf("Expected no errorHandler, got %v", err))
	}
	suite.Equal(http.StatusOK, resp.StatusCode)

	decoder := json.NewDecoder(resp.Body)
	var data dto.Orders
	_ = decoder.Decode(&data)

	suite.Equal(data.Count, uint(1))
}

func (suite *TestSuite) TestGetOrderByID() {
	utils.CreateOrderInDB(&suite.OrderModel, suite.Server.DB)

	resp, err := http.Get(fmt.Sprintf("%s/orders/%s", suite.Server.URL, suite.OrderModel.ID.String())) //nolint:bodyclose
	if err != nil {
		suite.Fail(fmt.Sprintf("Expected no errorHandler, got %v", err))
	}
	suite.Equal(http.StatusOK, resp.StatusCode)

	decoder := json.NewDecoder(resp.Body)
	var data dto.Order
	_ = decoder.Decode(&data)

	suite.Equal(data.OrderID, suite.OrderModel.ID)
}
func (suite *TestSuite) TestNotExistOrderID() {
	resp, err := http.Get(fmt.Sprintf("%s/orders/%s", suite.Server.URL, suite.OrderModel.ID.String())) //nolint:bodyclose
	if err != nil {
		suite.Fail(fmt.Sprintf("Expected no errorHandler, got %v", err))
	}
	suite.Equal(http.StatusNotFound, resp.StatusCode)
}
