package main

import (
	"github.com/MikhailGulkin/CleanGolangOrderApp/pkg/env"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/infrastructure/di"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/presentation/graph"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/presentation/grpc"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

func init() {
	err := godotenv.Load(env.GetEnv("CONFIG_PATH", "./configs/app/.env"))
	if err != nil {
		return
	}
}

func main() {
	fx.New(
		di.Module,
		graph.Module,
		grpc.Module,
	).Run()
}
