package biz

import (
	"github.com/google/wire"

	"pangud.io/pangud/internal/biz/user"
)

var ProviderSet = wire.NewSet(user.ProviderSet)
