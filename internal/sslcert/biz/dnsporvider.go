package biz

import (
	"go.uber.org/zap"
	"github.com/pangud/pkg/types"
)

// DNSProvider DNS服务商
type DNSProvider struct {
	ID     uint8  `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Key    string `json:"key"`
	Secret string `json:"secret"`
}

// TableName 表名
func (p *DNSProvider) TableName() string {
	return "t_core_ssl_dns_providers"
}

// DNSProviderRepository DNS服务商存储库
type DNSProviderRepository interface {
	types.Repository[*DNSProvider]
}

// DNSProviderUsecase DNS服务商用例
type DNSProviderUsecase struct {
	log  *zap.Logger
	repo DNSProviderRepository
}

// NewDNSProviderUsecase 创建DNS服务商用例
func NewDNSProviderUsecase(repo DNSProviderRepository, log *zap.Logger) *DNSProviderUsecase {
	return &DNSProviderUsecase{
		repo: repo,
		log:  log,
	}
}
