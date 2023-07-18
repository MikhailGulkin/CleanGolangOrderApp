package api

import (
	agg "github.com/MikhailGulkin/simpleGoOrderApp/internal/domain/order/aggregate"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/domain/product/aggregate"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/db/models"
	conftest2 "github.com/MikhailGulkin/simpleGoOrderApp/tests/integration/api/conftest"
	u "github.com/MikhailGulkin/simpleGoOrderApp/tests/integration/api/order/utils"
	"github.com/MikhailGulkin/simpleGoOrderApp/tests/integration/api/product/utils"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	Server         conftest2.Server
	ProductEntity  aggregate.Product
	ProductModel   models.Product
	OrderAggregate agg.Order
	OrderModel     models.Order
}

func (suite *TestSuite) SetupSuite() {
	server := conftest2.StartServer()
	defer server.Done()
	suite.Server = server
	suite.ProductEntity = utils.CreateValidProductEntity()
	suite.ProductModel = utils.CreateValidProductModel(suite.ProductEntity)
	suite.OrderAggregate = u.CreateValidOrderEntity()
	suite.OrderModel = u.CreateValidOrderModel(suite.OrderAggregate)
}
func (suite *TestSuite) TearDownTest() {
	conftest2.CleanTables(suite.Server.DB)
}
