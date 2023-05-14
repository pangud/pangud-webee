package data

import (
	"github.com/pangud/pangud/internal/core/biz"
	"go.uber.org/zap"
)

func (d *Data) Migrate() {
	if d.config.Application.IsMaster {
		//migrate user etc...
		err := d.db.AutoMigrate(&biz.User{}, biz.Endpoint{})
		if err != nil {
			d.log.Fatal("migrate user", zap.Error(err))
		}

	}
	d.log.Info("migrate database")
	// if err != nil {
	// 	d.log.Error("migrate database", zap.Error(err))
	// }
}
