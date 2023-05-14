package server

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/pangud/pangud/api"
	"github.com/pangud/pangud/internal/sslcert/resource"

	coreresource "github.com/pangud/pangud/internal/core/resource"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/pangud/pangud/internal/conf"
)

var ProviderSet = wire.NewSet(NewServer)

type Server struct {
	cfg        *conf.Bootstrap
	engine     *gin.Engine
	sslCertAPI *resource.SSLCertAPI
	coreAPI    *coreresource.CoreAPI
}

func NewServer(cfg *conf.Bootstrap, engine *gin.Engine, sslCertAPI *resource.SSLCertAPI, coreAPI *coreresource.CoreAPI) *Server {
	return &Server{
		cfg:        cfg,
		engine:     engine,
		sslCertAPI: sslCertAPI,
		coreAPI:    coreAPI,
	}
}

func (s *Server) Run() {
	//register routes
	// s.accountAPI.Register()
	s.sslCertAPI.Register()
	s.coreAPI.Register()
	api.SwaggerInfo.BasePath = "/api/v1"
	s.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	s.engine.Run(s.cfg.Server.Addr)
}
