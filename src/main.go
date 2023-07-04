package main

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/di"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api"
	"go.uber.org/fx"
)

//	type Config struct {
//		db.ConfigDB `toml:"db"`
//		API         struct {
//			Host string
//			Port int
//		} `toml:"api"`
//	}

func main() {
	fx.New(
		api.Module,
		db.Module,
		di.Module,
		fx.Invoke(api.Start),
	).Run()
	//err := providers.Provide(engine.Module).Start(context.Background())
	//if err != nil {
	//	return
	//}
	//var conf Config
	//config.LoadConfig(&conf, "")
	//fmt.Println(conf.API)
	//conf := db.NeWDBConfig()
	//connection := db.GetConnection(conf)
	//repo := dao.BuildProductDAO(connection)
	//controller := controllers.CreateProduct{Interactor: command.CreateProductImpl{ProductDAO: repo}}
	//
	//r := gin.Default()
	//r.POST("/create-product", controller.CreateProduct)
	//
	//err := r.Run("127.0.0.1:8000")
	//if err != nil {
	//	return
	//}

}
