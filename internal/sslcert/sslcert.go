package sslcert

import (
	"github.com/google/wire"

	"github.com/pangud/internal/sslcert/biz"
	"github.com/pangud/internal/sslcert/data"
)

var ProviderSet = wire.NewSet(biz.ProviderSet, data.ProviderSet)
