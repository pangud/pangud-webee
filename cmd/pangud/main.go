package main

import (
	"fmt"
	"go.uber.org/zap"
	"log"
	"pangud.io/pangud/internal/data"
	"pangud.io/pangud/internal/server"

	//"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"

	"pangud.io/pangud/internal/conf"
	log2 "pangud.io/pangud/internal/log"
)

func main() {
	fmt.Println("Feiku panel")
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
