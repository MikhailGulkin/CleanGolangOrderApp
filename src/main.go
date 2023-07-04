package main

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api"
	"go.uber.org/fx"
)

func main() {
	fx.New(api.Module).Run()
}
