package data

import (
	"context"
	"go.uber.org/zap"
	
	"pangud.io/pangud/internal/pkg/data"
	"pangud.io/pangud/internal/sslcert/biz"
)

type dnsProviderRepository struct {
	log  *zap.Logger
	data *data.Data
}

func (d dnsProviderRepository) Save(ctx context.Context, model *biz.DNSProvider) error {
	//*biz.DNSProviderODO implement me
	panic("implement me")
}

func (d dnsProviderRepository) FindOne(ctx context.Context, id uint32) (*biz.DNSProvider, error) {
	//*biz.DNSProviderODO implement me
	panic("implement me")
}

func (d dnsProviderRepository) Remove(ctx context.Context, model *biz.DNSProvider) error {
	//*biz.DNSProviderODO implement me
	panic("implement me")
}

// NewDNSProviderRepository 创建DNS服务商存储库
func NewDNSProviderRepository(data *data.Data, log *zap.Logger) biz.DNSProviderRepository {
	return &dnsProviderRepository{log: log, data: data}
}
