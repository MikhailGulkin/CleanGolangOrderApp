package di

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/application"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/infrastructure/db"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/infrastructure/minio"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	db.NewOutbox,
	db.NewConfig,
	db.NewConnection,
	db.NewEventStore,
	db.NewUoWManager,
	minio.NewMinio,
	application.NewCustomerServices,
)
