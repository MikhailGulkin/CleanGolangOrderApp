package conftest

import (
	"fmt"
	load "github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/config"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/di"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/config"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/controllers/routes"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/engine"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/providers"
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
	fmt.Println(conf.FullDNS())
	return conf
}

var ModuleEngine = fx.Provide(
	engine.NewBaseGroup,
	NewRequestHandler,
)
var ModuleConfig = fx.Provide(
	NewConfig,
	providers.NewDBConfig,
	providers.NewAPIConfig,
)
var Module = fx.Options(
	ModuleConfig,
	routes.Module,
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
