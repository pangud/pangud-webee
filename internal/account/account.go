package account

import (
	"github.com/google/wire"

	"pangud.io/pangud/internal/account/biz"
	"pangud.io/pangud/internal/account/data"
	"pangud.io/pangud/internal/account/resource"
)

var ProviderSet = wire.NewSet(biz.ProviderSet, data.ProviderSet, resource.ProviderSet)
