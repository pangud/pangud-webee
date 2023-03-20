package conf

type Bootstrap struct {
	Server  *conf.Server
	Data    *conf.Data
	Logger  *conf.Logger
	Workdir string
	Docker  *conf.Docker
}
