package conf

import (
	"github.com/pangud/pangud/pkg/conf"
)

type Bootstrap struct {
	Application *conf.Application `yaml:"app"`
	Server      *conf.Server
	Data        *conf.Data
	Logger      *conf.Logger
	Docker      *conf.Docker
}
