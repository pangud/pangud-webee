package data

import (
	"context"
	"path/filepath"

	"github.com/google/wire"
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/pangud/pangud/internal/conf"
	"github.com/pangud/pangud/pkg/log"
	"github.com/pangud/pangud/pkg/tx"
)

var ProviderSet = wire.NewSet(NewData, NewTransaction)

type Data struct {
	config *conf.Bootstrap
	// 通过DB(ctx)获取 以支持事务
	db  *gorm.DB
	log *zap.Logger
}

// 事务key
type contextTxKey struct{}

// NewTransaction new a transaction
func NewTransaction(data *Data) tx.Transaction {
	return data
}

// Execute 执行事务
func (d *Data) Execute(ctx context.Context, fn func(ctx context.Context) error) error {
	return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		d.log.Sugar().Debugf("start transaction")
		defer d.log.Sugar().Debugf("end transaction")
		ctx = context.WithValue(ctx, contextTxKey{}, tx)
		return fn(ctx)
	})
}

func (d *Data) DB(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(contextTxKey{}).(*gorm.DB); ok {
		return tx
	}
	return d.db.WithContext(ctx)
}

//func (d *Data) Rdb() *redis.Client {
//	return d.rdb
//}

// NewData new a data and return.
func NewData(cfg *conf.Bootstrap, logger *zap.Logger) (*Data, func(), error) {

	dsn := filepath.Join(cfg.Application.Workdir, "data/.pangud.sdb")

	logger.Sugar().Debug("dsn: ", dsn)

	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: log.NewGormLogger(cfg.Logger)})

	if err != nil {
		logger.Sugar().Fatalf("db connect error: %s", err)
	}

	//dbpath := filepath.Join(cfg.Workdir, "data/.pangud.bdb")
	//bdb, err := bbolt.Open(dbpath, 0666, nil)
	//if err != nil {
	//	logger.Sugar().Fatalf("db connect error: %s", err)
	//	//return err
	//}

	cleanup := func() {
		logger.Sugar().Info("closing the data resources")
		if conn, err := db.DB(); err == nil {
			conn.Close()
		}
		//bdb.Close()
	}

	return &Data{config: cfg, db: db, log: logger}, cleanup, nil
}
