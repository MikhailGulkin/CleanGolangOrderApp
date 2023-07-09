package main

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/di"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		di.Module,
		api.Module,
	).Run()
}
