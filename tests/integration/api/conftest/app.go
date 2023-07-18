package conftest

import (
	"context"
	"fmt"
	load "github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/config"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/db"
	dbFactory "github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/di/factories/db"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/di/factories/interactors"
	logger2 "github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/di/factories/logger"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/di/factories/mediator"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/logger"
	api "github.com/MikhailGulkin/simpleGoOrderApp/internal/presentation/api/config"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/presentation/api/controllers/routes"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/presentation/api/engine"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/presentation/api/providers/controllers"
	middleware2 "github.com/MikhailGulkin/simpleGoOrderApp/internal/presentation/api/providers/middleware"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/presentation/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"os"
)

func NewRequestHandler() engine.RequestHandler {
	gin.SetMode(gin.ReleaseMode)
	newEngine := gin.New()
	return engine.RequestHandler{Gin: newEngine}
}
func createConfFabric(dbPort int) func() config.Config {
	return func() config.Config {
		var conf config.Config
		load.LoadConfig(&conf, os.Getenv("PROJECT_PATH"), "./config/test.toml")
		conf.DBConfig.Port = dbPort
		return conf
	}
}

var ModuleEngine = fx.Provide(
	NewRequestHandler,
)
var ModuleConfig = fx.Provide(
	config.NewDBConfig,
	config.NewAPIConfig,
	config.NewLoggerConfig,
)
var DiModule = fx.Options(
	dbFactory.Module,
	interactors.Module,
	logger2.Module,
	mediator.Module,
)

var Module = fx.Options(
	ModuleConfig,
	DiModule,
	controllers.Module,
	middleware2.Module,
	ModuleEngine,
	fx.Invoke(Start),
)

type Server struct {
	*fx.App
	URL string
	*gorm.DB
}

func StartServer(dbPort int) Server {
	fabric := createConfFabric(dbPort)
	conf := fabric()

	app := fx.New(
		fx.NopLogger,
		Module,
		fx.Provide(createConfFabric(dbPort)),
	)
	go func() {
		app.Run()
	}()
	conn := db.BuildConnection(logger.Logger{}, conf.DBConfig)

	return Server{App: app, URL: setupBaseURL(conf.APIConfig), DB: conn}
}
func Start(
	lifecycle fx.Lifecycle,
	router engine.RequestHandler,
	config api.APIConfig,
	routers routes.Routes, //nolint:all
) {
	routers.Setup()

	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				go func() {
					err := router.Gin.Run(fmt.Sprintf("%s:%d", config.Host, config.Port))
					if err != nil {
						panic(err)
					}
				}()
				return nil
			},
		},
	)
}
