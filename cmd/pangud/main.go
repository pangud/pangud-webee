package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	conf2 "github.com/pangud/pangud/internal/conf"
	"github.com/pangud/pangud/internal/pkg/data"
	"github.com/pangud/pangud/internal/server"
	"github.com/pangud/pangud/pkg/conf"
	log2 "github.com/pangud/pangud/pkg/log"
)

// @title           PangudOS API
// @version         1.0
// @description     PANGUD OS API.
// @termsOfService  https://pangud.org

// @contact.name   服务支持
// @contact.url    https://pangud.org
// @contact.email  dev_support@gail.com

// @license.name  AGPL-3.0
// @license.url   https://www.gnu.org/licenses/agpl-3.0.en.html

// @host      localhost:2345
// @BasePath  /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	fmt.Println("pangud")
	var bc = conf2.Bootstrap{}
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

func newApp(data *data.Data, server *server.Server, cfg *conf2.Bootstrap, logger *zap.Logger) *App {
	fmt.Println("Pangu OS")
	return &App{
		data:   data,
		server: server,
		conf:   cfg,
		logger: logger,
	}
}
