package api

import (
	"fmt"
	"github.com/MikhailGulkin/simpleGoOrderApp/tests/integration/api/order/utils"
	"net/http"
)

func (suite *TestSuite) TestSuccessOrderCreate() {
	utils.CreateAddressInDB(&suite.OrderModel, suite.Server.DB)
	utils.CreateProductInDB(&suite.OrderModel, suite.Server.DB)
	utils.CreateClientInDB(&suite.OrderModel, suite.Server.DB)

	resp, err := http.Post( //nolint:bodyclose
		fmt.Sprintf("%s/orders", suite.Server.URL),
		"application/json",
		utils.CreateValidByteOrder(suite.OrderModel),
	)

	if err != nil {
		suite.Fail(fmt.Sprintf("Expected no errorHandler, got %v", err))
	}

	suite.Equal(http.StatusNoContent, resp.StatusCode)
	suite.Equal(resp.ContentLength, ZeroInt64)
}
