package conftest

import (
	"context"
	"fmt"
	load "github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/config"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db"
	commandbus "github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/di/factories/commandBus"
	dbFactory "github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/di/factories/db"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/di/factories/interactors"
	logger2 "github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/di/factories/logger"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/logger"
	api "github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/config"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/controllers/routes"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/engine"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/providers/controllers"
	middleware2 "github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/providers/middleware"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"os"
	"strconv"
)

func NewRequestHandler() engine.RequestHandler {
	gin.SetMode(gin.ReleaseMode)
	newEngine := gin.New()
	return engine.RequestHandler{Gin: newEngine}
}

func NewConfig() config.Config {
	var conf config.Config
	load.LoadConfig(&conf, os.Getenv("PROJECT_PATH"), "./config/test.toml")
	return conf
}

var ModuleEngine = fx.Provide(
	NewRequestHandler,
)
var ModuleConfig = fx.Provide(
	NewConfig,
	config.NewDBConfig,
	config.NewAPIConfig,
	config.NewLoggerConfig,
)
var DiModule = fx.Options(
	dbFactory.Module,
	interactors.Module,
	logger2.Module,
	commandbus.Module,
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

func StartServer() Server {
	conf := NewConfig()
	app := fx.New(
		fx.NopLogger,
		Module,
	)
	go func() {
		app.Run()
	}()
	waitForServer(strconv.Itoa(conf.APIConfig.Port))
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
