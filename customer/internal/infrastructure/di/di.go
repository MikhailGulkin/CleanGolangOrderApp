package di

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/application"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/application/persistence"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/aggregate"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/infrastructure/db"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/infrastructure/minio"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	fx.Annotate(
		db.NewOutbox,
		fx.As(new(persistence.Outbox)),
	),
	fx.Annotate(
		db.NewEventStore,
		fx.As(new(aggregate.EventStore)),
	),
	fx.Annotate(
		minio.NewMinio,
		fx.As(new(persistence.Bucket)),
	),
	fx.Annotate(
		db.NewUoWManager,
		fx.As(new(persistence.UoWManager)),
	),
	db.NewConfig,
	db.NewConnection,
	application.NewCustomerServices,
)
