package server

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	"pangud.io/pangud/internal/api"
	"pangud.io/pangud/internal/conf"
)

var ProviderSet = wire.NewSet(NewServer)

type Server struct {
	cfg          *conf.Bootstrap
	engine       *gin.Engine
	userResource *api.UserResource
}

func NewServer(cfg *conf.Bootstrap, engine *gin.Engine, userResource *api.UserResource) *Server {
	return &Server{
		cfg:          cfg,
		engine:       engine,
		userResource: userResource,
	}
}

func (s *Server) Run() {
	//register routes
	apiv1 := s.engine.Group("/api/v1")
	apiv1.Group("users").GET("", s.userResource.List)
	s.engine.Run(s.cfg.Server.Addr)
}
