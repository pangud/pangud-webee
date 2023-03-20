package resource

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewSSLCertPI, NewDNSProviderResource)

type SSLCertAPI struct {
	engine              *gin.Engine
	dnsProviderResource *DNSProviderResource
}

// NewSSLCertPI new an account api and return.
func NewSSLCertPI(engine *gin.Engine, dnsProviderResource *DNSProviderResource) *SSLCertAPI {
	return &SSLCertAPI{
		engine:              engine,
		dnsProviderResource: dnsProviderResource,
	}
}

func (api *SSLCertAPI) Register() {
	apiv1 := api.engine.Group("/api/v1")
	group := apiv1.Group("ssl_certs")

	group.GET("dns_providers", api.dnsProviderResource.List)
}
