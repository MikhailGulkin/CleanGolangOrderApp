package api

import (
	"context"
	"fmt"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/config"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/engine"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/routes"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewConfig),
	routes.Module,
	engine.Module,
)

type Config struct {
	db.ConfigDB
	APIConfig struct {
		Host string
		Port int
	} `toml:"api"`
}

func NewConfig() Config {
	var conf Config
	config.LoadConfig(&conf, "")
	return conf
}

func Start(lifecycle fx.Lifecycle, router engine.RequestHandler, config Config, routes routes.Routes) {
	routes.Setup()
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				fmt.Printf("Starting application in :%d", config.APIConfig.Port)
				go router.Gin.Run(fmt.Sprintf("%s:%d", config.APIConfig.Host, config.APIConfig.Port))
				return nil
			},
			OnStop: func(context.Context) error {
				fmt.Println("Stopping application")
				return nil
			},
		},
	)
}
