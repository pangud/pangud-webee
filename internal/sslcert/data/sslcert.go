package data

import (
	"context"
	"go.uber.org/zap"
	"pangud.io/pangud/internal/pkg/data"
	"pangud.io/pangud/internal/sslcert/biz"
)

type sslcertRepository struct {
	log  *zap.Logger
	data *data.Data
}

// NewSSLCertRepository 创建证书存储库
func NewSSLCertRepository(data *data.Data, log *zap.Logger) biz.SSLCertRepository {
	return &sslcertRepository{log: log, data: data}
}

func (s sslcertRepository) Save(ctx context.Context, model *biz.SSLCert) error {
	//TODO implement me
	panic("implement me")
}

func (s sslcertRepository) FindOne(ctx context.Context, id uint32) (*biz.SSLCert, error) {
	//TODO implement me
	panic("implement me")
}

func (s sslcertRepository) Remove(ctx context.Context, model *biz.SSLCert) error {
	//TODO implement me
	panic("implement me")
}
