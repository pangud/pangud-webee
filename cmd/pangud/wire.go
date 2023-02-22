//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go.uber.org/zap"

	"pangud.io/pangud/internal/biz"
	"pangud.io/pangud/internal/data"
	"pangud.io/pangud/internal/interface/restful"
	"pangud.io/pangud/internal/server"
	"pangud.io/pangud/pkg/conf"
)

// App is the main application
type App struct {
	server *server.Server
	router *gin.Engine
	conf   *conf.Bootstrap
	data   *data.Data
	logger *zap.Logger
}

func (a *App) Run() {
	a.data.Migrate()
	a.server.Run()
}
func wireApp(cfg *conf.Bootstrap, engine *gin.Engine, logger *zap.Logger) (*App, func(), error) {
	panic(wire.Build(data.ProviderSet, restful.ProviderSet, server.ProviderSet, biz.ProviderSet, newApp))
}
