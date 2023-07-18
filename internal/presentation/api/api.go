package api

import (
	"context"
	"fmt"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/logger"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/presentation/api/config"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/presentation/api/controllers/routes"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/presentation/api/engine"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/presentation/api/providers"
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
) {
	routers.Setup()

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
