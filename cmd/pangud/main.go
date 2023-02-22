package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"pangud.io/pangud/internal/data"
	"pangud.io/pangud/internal/server"
	"pangud.io/pangud/pkg/conf"
	log2 "pangud.io/pangud/pkg/log"
)

func main() {
	fmt.Println("pangud server")
	var bc = conf.Bootstrap{}
	err := conf.Load("./configs/config.yaml", &bc)
	if err != nil {
		log.Fatalln(err)
	}
	router := gin.Default()
	//router.Use(cors.Default())

	logger := log2.New(bc.Logger, "core.log")
	app, cleanup, err := wireApp(&bc, router, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	logger.Sugar().Info(zap.Any("app", app))

	//migrate

	app.Run()
}

func newApp(data *data.Data, server *server.Server, cfg *conf.Bootstrap, logger *zap.Logger) *App {
	fmt.Println("Pangu OS")
	return &App{
		data:   data,
		server: server,
		conf:   cfg,
		logger: logger,
	}
}
