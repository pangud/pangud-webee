package biz

import (
	"go.uber.org/zap"
	"github.com/pangud/pangud/pkg/types"
)

// SSLCert 证书
type SSLCert struct {
	ID          uint32 `json:"id"`
	Host        string `json:"host"` // 域名
	PrivateKey  string `json:"private_key"`
	Certificate string `json:"certificate"`
	Addition    string `json:"addition"`
}

func (c *SSLCert) TableName() string {
	return "t_core_ssl_certs"
}

// SSLCertRepository 证书存储库
type SSLCertRepository interface {
	types.Repository[*SSLCert]
}

// SSLCertUsecase 证书用例
type SSLCertUsecase struct {
	log             *zap.Logger
	certRepo        SSLCertRepository
	dnsProviderRepo DNSProviderRepository
}

// NewSSLCertUsecase 创建证书用例
func NewSSLCertUsecase(repo SSLCertRepository, dnsProviderRepo DNSProviderRepository, log *zap.Logger) *SSLCertUsecase {
	return &SSLCertUsecase{
		certRepo:        repo,
		dnsProviderRepo: dnsProviderRepo,
		log:             log,
	}
}
