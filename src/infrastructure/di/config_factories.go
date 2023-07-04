package di

import "C"
import (
	c "github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/command"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/command"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/persistence"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/config"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db"
	"go.uber.org/fx"
)

func NewDBConfig() db.ConfigDB {
	var conf db.ConfigDB
	config.LoadConfig(&conf, "db")

	return conf
}
func NewCreateProductImpl(dao persistence.ProductDAO) command.CreateProduct {
	product := c.CreateProductImpl{
		ProductDAO: dao,
	}
	return &product
}

var Module = fx.Provide(
	NewDBConfig,
	NewCreateProductImpl,
)
