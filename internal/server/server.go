package server

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	"pangud.io/pangud/pkg/conf"

	account "pangud.io/pangud/internal/account/resource"
)

var ProviderSet = wire.NewSet(NewServer)

type Server struct {
	cfg        *conf.Bootstrap
	engine     *gin.Engine
	accountAPI *account.AccountAPI
}

func NewServer(cfg *conf.Bootstrap, engine *gin.Engine, accountAPI *account.AccountAPI) *Server {
	return &Server{
		cfg:        cfg,
		engine:     engine,
		accountAPI: accountAPI,
	}
}

func (s *Server) Run() {
	//register routes
	s.accountAPI.Register()
	s.engine.Run(s.cfg.Server.Addr)
}
