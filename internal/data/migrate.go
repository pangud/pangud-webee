package data

import (
	"go.uber.org/zap"

	"pangud.io/pangud/internal/apiserver/biz"
)

func (d *Data) Migrate() {
	d.log.Info("migrate database")
	err := d.db.AutoMigrate(&biz.User{})
	if err != nil {
		d.log.Error("migrate database", zap.Error(err))
	}
}
