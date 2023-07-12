package conftest

import (
	load "github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/config"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/di"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/controllers/routes"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/engine"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/middleware"
	config2 "github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/config"
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

func NewConfig() config2.Config {
	var conf config2.Config
	load.LoadConfig(&conf, os.Getenv("PROJECT_PATH"), "./config/test.toml")
	return conf
}

var ModuleEngine = fx.Provide(
	NewRequestHandler,
)
var ModuleConfig = fx.Provide(
	NewConfig,
	config2.NewDBConfig,
	config2.NewAPIConfig,
)
var Module = fx.Options(
	ModuleConfig,
	routes.Module,
	middleware.Module,
	ModuleEngine,
	di.Module,
	fx.Invoke(api.Start),
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
	conn := db.BuildConnection(conf.DBConfig)

	return Server{App: app, URL: setupBaseURL(conf.APIConfig), DB: conn}
}
