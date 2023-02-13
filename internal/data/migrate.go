package data

import (
	"go.uber.org/zap"

	"pangud.io/pangud/internal/biz/user"
)

func (d *Data) Migrate() {
	d.log.Info("migrate database")
	err := d.db.AutoMigrate(&user.User{})
	if err != nil {
		d.log.Error("migrate database", zap.Error(err))
	}
}