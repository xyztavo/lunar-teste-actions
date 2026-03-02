package utils

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewValidator),
	fx.Provide(NewUtils),
)
