package sslcert

import (
	"github.com/google/wire"

	"github.com/pangud/pangud/internal/sslcert/biz"
	"github.com/pangud/pangud/internal/sslcert/data"
	"github.com/pangud/pangud/internal/sslcert/resource"
)

var ProviderSet = wire.NewSet(biz.ProviderSet, data.ProviderSet, resource.ProviderSet)
