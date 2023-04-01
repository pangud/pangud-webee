package container

import (
	"github.com/google/wire"

	"github.com/pangud/pangud/internal/container/internal/biz"
	"github.com/pangud/pangud/internal/container/internal/resource"
)

var ProviderSet = wire.NewSet(biz.ProviderSet, resource.ProviderSet)
