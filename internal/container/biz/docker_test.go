package biz

import (
	"fmt"
	"log"
	"os"
	"testing"

	conf2 "github.com/pangud/pangud/internal/conf"
	"github.com/pangud/pangud/pkg/conf"
	log2 "github.com/pangud/pangud/pkg/log"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	if code != 0 {
		os.Exit(code)
	}
}

var imageUsecase *ImageUsecase

func setup() {
	fmt.Println("pangud server")
	var bc = conf2.Bootstrap{}
	err := conf.Load("/Users/liwei/MyWorkspace/pangud/pangud/configs/config.yaml", &bc)
	if err != nil {
		log.Fatalln(err)
	}
	//router.Use(cors.Default())

	logger := log2.New(bc.Logger, "core.log")

	cli := NewDockerClient(&bc, logger)
	imageUsecase = NewImageUsecase(cli, logger)
}
