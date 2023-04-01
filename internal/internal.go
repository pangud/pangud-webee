package internal

import (
	"github.com/google/wire"

	"github.com/pangud/pangud/internal/container"
	"github.com/pangud/pangud/internal/pkg/data"
	"github.com/pangud/pangud/internal/server"
	"github.com/pangud/pangud/internal/sslcert"

	"github.com/pangud/pangud/internal/sslcert/resource"
)

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(data.ProviderSet, sslcert.ProviderSet, container.ProviderSet, resource.ProviderSet, server.ProviderSet)
