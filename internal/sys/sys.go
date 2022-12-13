package sys

import (
	"github.com/google/wire"

	"pangud.io/pangud/internal/sys/user"
)

var ProviderSet = wire.NewSet(user.ProviderSet)
