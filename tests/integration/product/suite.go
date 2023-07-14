package product

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/product/aggregate"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/models"
	"github.com/MikhailGulkin/simpleGoOrderApp/tests/integration/conftest"
	"github.com/MikhailGulkin/simpleGoOrderApp/tests/integration/product/utils"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	Server        conftest.Server
	ProductEntity aggregate.Product
	ProductModel  models.Product
}

func (suite *TestSuite) SetupSuite() {
	server := conftest.StartServer()
	defer server.Done()
	suite.Server = server
	suite.ProductEntity = utils.CreateValidProductEntity()
	suite.ProductModel = utils.CreateValidProductModel(suite.ProductEntity)
}
func (suite *TestSuite) TearDownTest() {
	conftest.CleanTables(suite.Server.DB)
}
