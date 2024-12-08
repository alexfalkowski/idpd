package pipeline

import (
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	fx.Provide(NewService),
	fx.Provide(NewRepository),
	fx.Provide(NewCommand),
)
