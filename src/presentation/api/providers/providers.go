package providers

import "go.uber.org/fx"

var Module = fx.Provide(
	NewConfig,
	NewDBConfig,
	NewAPIConfig,
)
