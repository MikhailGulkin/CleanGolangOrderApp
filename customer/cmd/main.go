package main

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/infrastructure/di"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/presentation/api"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/presentation/graph"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/presentation/grpc"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		di.Module,
		graph.Module,
		grpc.Module,
		api.Module,
	).Run()
}
