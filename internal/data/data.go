package data

import (
	"context"
	"path/filepath"
	"time"

	"github.com/go-redis/redis/v8"
	bolt "go.etcd.io/bbolt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"pangud.io/pangud/internal/conf"
	"pangud.io/pangud/internal/log"
	"pangud.io/pangud/internal/tx"
)

type Data struct {
	// 通过DB(ctx)获取 以支持事务
	db     *gorm.DB
	rdb    *redis.Client
	boltDB *bolt.DB
	log    *zap.Logger
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
	return d.db
}

func (d *Data) Rdb() *redis.Client {
	return d.rdb
}

// NewData new a data and return.
func NewData(cfg *conf.Bootstrap, logger *zap.Logger) (*Data, func(), error) {

	dsn := cfg.Data.Database.DSN

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: log.NewGormLogger(cfg.Logger)})

	if err != nil {
		logger.Sugar().Fatalf("db connect error: %s", err)
	}

	rdbCfg := cfg.Data.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:         rdbCfg.Addr,
		Password:     rdbCfg.Password,      // no password set
		DB:           int(rdbCfg.Database), // use default DB
		IdleTimeout:  time.Duration(rdbCfg.IdleTimeout) * time.Second,
		ReadTimeout:  rdbCfg.ReadTimeout,
		WriteTimeout: rdbCfg.WriteTimeout,
		MinIdleConns: int(rdbCfg.MaxIdle),
	})

	dbpath := filepath.Join(cfg.Workdir, "pangud.db")
	bdb, err := bolt.Open(dbpath, 0666, nil)
	if err != nil {
		logger.Sugar().Fatalf("db connect error: %s", err)
		//return err
	}

	cleanup := func() {
		logger.Sugar().Info("closing the data resources")
		if conn, err := db.DB(); err == nil {
			conn.Close()
		}
		rdb.Close()
		bdb.Close()
	}

	return &Data{db: db, rdb: rdb, boltDB: bdb, log: logger}, cleanup, nil
}
