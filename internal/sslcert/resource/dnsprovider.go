package resource

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/pangud/pangud/internal/sslcert/biz"
)

type DNSProviderResource struct {
	log                *zap.Logger
	dnsProviderUsecase *biz.DNSProviderUsecase
}

func NewDNSProviderResource(logger *zap.Logger, dnsProviderUsecase *biz.DNSProviderUsecase) *DNSProviderResource {
	return &DNSProviderResource{
		log:                logger,
		dnsProviderUsecase: dnsProviderUsecase,
	}
}

// List 查询DNS提供商列表
// @Summary 查询DNS提供商列表
// @Schemes
// @Description 查询DNS提供商列表
// @Tags SSL证书
// @Accept json
// @Produce json
// @Success 200 {string} string	"ok"
// @Router /ssl_certs/dns_providers [get]
// @Security ApiKeyAuth
func (r *DNSProviderResource) List(ctx *gin.Context) {
	r.log.Debug("list dns providers")
	//r.userUsecase.List(ctx)
}
