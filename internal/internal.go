package internal

import (
	"github.com/google/wire"

	"github.com/pangud/pangud/internal/container"
	"github.com/pangud/pangud/internal/core"
	"github.com/pangud/pangud/internal/pkg/data"
	"github.com/pangud/pangud/internal/server"
	"github.com/pangud/pangud/internal/sslcert"
)

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(data.ProviderSet, core.ProviderSet, sslcert.ProviderSet, container.ProviderSet, server.ProviderSet)
