package biz

import (
	"github.com/docker/docker/client"
	"github.com/google/wire"
	"go.uber.org/zap"

	"github.com/pangud/pangud/pkg/conf"
)

var ProviderSet = wire.NewSet(NewDockerClient)

func NewDockerClient(cfg *conf.Bootstrap, logger *zap.Logger) *client.Client {
	cli, err := client.NewClientWithOpts(client.WithHost(cfg.Docker.Host), client.WithVersion("1.39"))
	if err != nil {
		panic(err)
	}
	return cli
}
