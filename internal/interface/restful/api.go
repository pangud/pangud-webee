package restful

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewUserResource)
