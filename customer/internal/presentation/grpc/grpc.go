package grpc

import (
	"context"
	"fmt"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		NewGrpcServer,
		NewCustomerGrpcService,
	),
	fx.Invoke(Start),
)

func Start(lifecycle fx.Lifecycle, server *GrpcServer) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				go func() {
					defer func() {
						if r := recover(); r != nil {
							fmt.Printf("Recovered when boot grpc server, r %s", r)
						}
					}()
					err := server.Run()
					if err != nil {
						panic(err)
					}
				}()
				return nil
			},
			OnStop: func(context.Context) error {
				return nil
			},
		},
	)
}
