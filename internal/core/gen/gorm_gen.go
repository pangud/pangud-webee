package main

import (
	"fmt"

	"github.com/pangud/pangud/internal/core/internal/biz"
	"gorm.io/gen"
)

func main() {
	fmt.Printf("Gen data object.\n")
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/core/internal/data",
		Mode:    gen.WithDefaultQuery,
	})

	g.ApplyBasic(biz.User{}, biz.Endpoint{})

	// g.ApplyInterface(func(model.Method) {}, model.User{})
	// g.ApplyInterface(func(model.UserMethod) {}, model.User{})

	g.Execute()
}
