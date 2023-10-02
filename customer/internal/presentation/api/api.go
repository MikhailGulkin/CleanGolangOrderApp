package api

import (
	"context"
	"fmt"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		NewEngine,
		NewCustomerHandler,
		NewMiddlewares,
		NewFiberGroup,
		fx.Annotate(
			NewCustomerRouter,
			fx.As(new(Route)),
		),
	),
	fx.Invoke(Start),
)

func Start(
	lifecycle fx.Lifecycle,
	route Route,
	engine Engine,
) {
	engine.Setup()
	route.Setup()

	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				go func() {
					defer func() {
						if r := recover(); r != nil {
							fmt.Printf("Recovered when boot grpc server, r %s", r)
						}
					}()
					err := engine.Run()
					if err != nil {
						panic(err)
					}
				}()
				return nil
			},
			OnStop: func(context.Context) error {
				return engine.Shutdown()
			},
		},
	)
}
