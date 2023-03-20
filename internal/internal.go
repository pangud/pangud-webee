package internal

import (
	"github.com/google/wire"

	"pangud.io/pangud/internal/pkg/data"
	"pangud.io/pangud/internal/server"
	"pangud.io/pangud/internal/sslcert"
)

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(data.ProviderSet, sslcert.ProviderSet, server.ProviderSet)
