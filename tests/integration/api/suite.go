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
	DBContainer    *conftest2.TestDatabase
}

func (suite *TestSuite) SetupSuite() {
	testDB := conftest2.NewTestDatabase(suite.T())
	//defer testDB.Close(suite.T())

	server := conftest2.StartServer(testDB.Port(suite.T()))
	defer server.Done()
	suite.DBContainer = testDB
	suite.Server = server
	suite.ProductEntity = utils.CreateValidProductEntity()
	suite.ProductModel = utils.CreateValidProductModel(suite.ProductEntity)
	suite.OrderAggregate = u.CreateValidOrderEntity()
	suite.OrderModel = u.CreateValidOrderModel(suite.OrderAggregate)
}
func (suite *TestSuite) TearDownTest() {
	conftest2.CleanTables(suite.Server.DB)
}
