package di

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/application"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/infrastructure/db"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	db.NewConfig,
	db.NewConnection,
	db.NewEventStore,
	db.NewUoWManager,
	application.NewCustomerServices,
)
