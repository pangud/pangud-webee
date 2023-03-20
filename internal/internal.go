package internal

import (
	"github.com/google/wire"

	"github.com/pangud/internal/pkg/data"
	"github.com/pangud/internal/server"
	"github.com/pangud/internal/sslcert"
)

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(data.ProviderSet, sslcert.ProviderSet, server.ProviderSet)
