package api

import (
	"context"
	"fmt"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/logger"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/presentation/api/config"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/presentation/api/controllers/routes"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/presentation/api/engine"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/presentation/api/prometheus"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/presentation/api/providers"
	"go.uber.org/fx"
)

var Module = fx.Options(
	providers.Module,
	fx.Invoke(Start),
)

func Start(
	lifecycle fx.Lifecycle,
	router engine.RequestHandler,
	config config.APIConfig,
	logger logger.Logger,
	routers routes.Routes, //nolint:all
	prometheus prometheus.Prometheus,
) {
	routers.Setup()
	prometheus.Use(router.Gin)

	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				logger.Info(fmt.Sprintf("Starting application in :%d", config.Port))
				go func() {
					defer func() {
						if r := recover(); r != nil {
							logger.Info(fmt.Sprintf("Recovered when boot api server, r %s", r))
						}
					}()
					err := router.Gin.Run(fmt.Sprintf("%s:%d", config.Host, config.Port))
					if err != nil {
						panic(err)
					}
				}()
				return nil
			},
			OnStop: func(context.Context) error {
				logger.Info("Stopping api application")
				return nil
			},
		},
	)
}
