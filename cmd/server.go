package cmd

import (
	"github.com/alexfalkowski/go-service/compress"
	"github.com/alexfalkowski/go-service/crypto"
	"github.com/alexfalkowski/go-service/debug"
	"github.com/alexfalkowski/go-service/encoding"
	"github.com/alexfalkowski/go-service/feature"
	"github.com/alexfalkowski/go-service/runtime"
	"github.com/alexfalkowski/go-service/sync"
	"github.com/alexfalkowski/go-service/telemetry"
	"github.com/alexfalkowski/go-service/transport"
	"github.com/alexfalkowski/idpd/config"
	"github.com/alexfalkowski/idpd/pipeline"
	"github.com/alexfalkowski/idpd/server/health"
	v1 "github.com/alexfalkowski/idpd/server/v1"
	"github.com/alexfalkowski/idpd/token"
	"go.uber.org/fx"
)

// ServerOptions for cmd.
var ServerOptions = []fx.Option{
	sync.Module, compress.Module, encoding.Module,
	runtime.Module, debug.Module, feature.Module,
	telemetry.Module, transport.Module,
	crypto.Module, token.Module,
	config.Module, health.Module,
	pipeline.Module, v1.Module, Module,
}
