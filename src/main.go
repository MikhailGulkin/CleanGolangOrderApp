package main

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/di"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/config"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/relay"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		di.Module,
		config.Module,
		api.Module,
		relay.Module,
	).Run()
}
