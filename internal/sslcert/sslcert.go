package sslcert

import (
	"github.com/google/wire"

	"pangud.io/pangud/internal/sslcert/biz"
	"pangud.io/pangud/internal/sslcert/data"
)

var ProviderSet = wire.NewSet(biz.ProviderSet, data.ProviderSet)
