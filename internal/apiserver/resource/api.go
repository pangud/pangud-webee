package resource

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewUserResource)