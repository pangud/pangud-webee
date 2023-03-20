package data

func (d *Data) Migrate() {
	d.log.Info("migrate database")
	// err := d.db.AutoMigrate(&accountBiz.User{})
	// if err != nil {
	// 	d.log.Error("migrate database", zap.Error(err))
	// }
}
