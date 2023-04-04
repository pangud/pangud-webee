package biz

import (
	"context"
	"time"

	"go.uber.org/zap"

	"github.com/pangud/pangud/internal/sslcert/internal/util"
	"github.com/pangud/pangud/pkg/types"
)

// SSLCert 证书
type SSLCert struct {
	ID          uint32    `json:"id"`
	Host        string    `json:"host"` // 域名
	PrivateKey  string    `json:"private_key"`
	Certificate string    `json:"certificate"`
	Addition    string    `json:"addition"`
	CreateTime  time.Time `gorm:"column:create_time;autoCreateTime"`
}

func (c *SSLCert) TableName() string {
	return "t_sslcert_certs"
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

func (s *SSLCertUsecase) Create(ctx context.Context, host, mail string) error {
	//todo find dnsprovider

	provider, err := s.dnsProviderRepo.FindOne(ctx, 1)
	//todo create cert

	if err != nil {
		return err
	}

	if pk, cert, err := util.GenCertByDnspod(host, mail, provider.Key, provider.Secret); err == nil {
		//save cert
		cert := &SSLCert{
			Host:        host,
			PrivateKey:  string(pk),
			Certificate: string(cert),
			CreateTime:  time.Now(),
		}
		err = s.certRepo.Save(ctx, cert)
		if err != nil {
			return err
		}

		return nil
	}
	return err
}
