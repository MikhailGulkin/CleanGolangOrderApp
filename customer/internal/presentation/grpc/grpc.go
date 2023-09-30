package grpc

import (
	"context"
	"fmt"
	"github.com/MikhailGulkin/CleanGolangOrderApp/pkg/env"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		NewGrpcServer,
		NewCustomerGrpcService,
	),
	fx.Invoke(Start),
)

func Start(lifecycle fx.Lifecycle, server *Server) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				go func() {
					defer func() {
						if r := recover(); r != nil {
							fmt.Printf("Recovered when boot grpc server, r %s", r)
						}
					}()
					err := server.Run(env.GetEnv("GRPC_PORT", "50052"))
					if err != nil {
						panic(err)
					}
				}()
				return nil
			},
			OnStop: func(context.Context) error {
				server.Down()
				return nil
			},
		},
	)
}
