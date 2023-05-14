package core

import (
	"context"

	"github.com/google/wire"
	"go.uber.org/zap"

	"github.com/pangud/pangud/internal/core/biz"
	"github.com/pangud/pangud/internal/core/data"
	"github.com/pangud/pangud/internal/core/resource"
	origindata "github.com/pangud/pangud/internal/pkg/data"
)

var ProviderSet = wire.NewSet(biz.ProviderSet, data.ProviderSet, resource.ProviderSet)

func Init(db *origindata.Data, log *zap.Logger) {
	data.SetDefault(db.DB(context.Background()))
	log.Info("core init success")
}
