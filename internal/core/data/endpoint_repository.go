package data

import (
	"context"

	"go.uber.org/zap"

	"github.com/pangud/pangud/internal/pkg/data"
)

type endpointRepository struct {
	data     *data.Data
	log      *zap.Logger
	endpoint endpoint
}

func NewEndpointRepository(data *data.Data, log *zap.Logger) *endpointRepository {
	return &endpointRepository{data: data, log: log, endpoint: newEndpoint(data.DB(context.Background()))}
}
