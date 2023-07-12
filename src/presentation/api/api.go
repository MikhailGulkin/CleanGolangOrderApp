package api

import (
	"context"
	"fmt"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/config"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/controllers/routes"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/engine"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/middleware"
	"go.uber.org/fx"
)

var Module = fx.Options(
	middleware.Module,
	routes.Module,
	engine.Module,
	fx.Invoke(Start),
)

func Start(
	lifecycle fx.Lifecycle,
	router engine.RequestHandler,
	config config.APIConfig,
	routers routes.Routes, //nolint:all
) {
	routers.Setup()

	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				fmt.Printf("Starting application in :%d", config.Port)
				go func() {
					err := router.Gin.Run(fmt.Sprintf("%s:%d", config.Host, config.Port))
					if err != nil {
						panic(err)
					}
				}()
				return nil
			},
			OnStop: func(context.Context) error {
				fmt.Println("Stopping application")
				return nil
			},
		},
	)
}
