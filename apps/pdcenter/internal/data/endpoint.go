package data

import (
	"context"

	"go.uber.org/zap"

	"github.com/pangud-apps/pdcenter/internal/biz"

	"github.com/pangud/pangud/internal/pkg/data"
)

type agentReadRepository struct {
	data *data.Data
	log  *zap.Logger
}

// NewAgentReadRepository new an agent read repository
func NewAgentReadRepository(data *data.Data,
	log *zap.Logger) biz.AgentReadRepository {
	return &agentReadRepository{
		data: data,
		log:  log,
	}
}

// FindOne find one
func (a agentReadRepository) FindOne(ctx context.Context, id uint32) (*biz.Agent, error) {
	var agent biz.Agent
	if err := a.data.DB(ctx).First(&agent, id).Error; err != nil {
		return nil, err
	}
	return &agent, nil
}
