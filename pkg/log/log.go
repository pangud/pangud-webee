package log

import (
	"io"
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	glogger "gorm.io/gorm/logger"

	"pangud.io/pangud/pkg/conf"
	"pangud.io/pangud/third_party/zapgorm2"
)

type WriteSyncer struct {
	io.Writer
}

func (ws WriteSyncer) Sync() error {
	return nil
}

func getWriteSyncer(cfg *conf.FileLogger, fileName string) zapcore.WriteSyncer {

	fileName = filepath.Join(cfg.Path, fileName)

	var ioWriter = &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    cfg.MaxSize,    // MB
		MaxBackups: cfg.MaxBackups, // number of backups
		MaxAge:     cfg.MaxAge,     //days
		LocalTime:  cfg.LocalTime,
		Compress:   cfg.Compress, // disabled by default
	}
	var sw = WriteSyncer{
		ioWriter,
	}
	return sw
}

func New(cfg *conf.Logger, modLogFile string) *zap.Logger {

	encoderCfg := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		TimeKey:        "time",
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"),
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	consoleDebugging := zapcore.Lock(os.Stdout)
	var fileLogLevel, consoleLogLevel zapcore.Level
	if cfg.File.Level == "info" {
		fileLogLevel = zap.InfoLevel
	} else if cfg.File.Level == "debug" {
		fileLogLevel = zap.DebugLevel
	} else {
		fileLogLevel = zap.InfoLevel
	}

	if cfg.Console.Level == "info" {
		consoleLogLevel = zap.InfoLevel
	} else if cfg.Console.Level == "debug" {
		consoleLogLevel = zap.DebugLevel
	} else {
		consoleLogLevel = zap.InfoLevel
	}
	var cores = make([]zapcore.Core, 0)
	cores = append(cores, zapcore.NewCore(zapcore.NewConsoleEncoder(encoderCfg), getWriteSyncer(cfg.File, cfg.File.Name), fileLogLevel))
	cores = append(cores, zapcore.NewCore(zapcore.NewConsoleEncoder(encoderCfg), consoleDebugging, consoleLogLevel))
	//isGorm := false

	if modLogFile != "" {
		cores = append(cores, zapcore.NewCore(zapcore.NewConsoleEncoder(encoderCfg),
			getWriteSyncer(cfg.File, modLogFile), fileLogLevel))
	}

	core := zapcore.NewTee(cores...)
	var zLogger *zap.Logger
	if consoleLogLevel == zap.DebugLevel || fileLogLevel == zap.DebugLevel {
		zLogger = zap.New(core, zap.AddCaller()).WithOptions()
	} else {
		zLogger = zap.New(core).WithOptions()
	}

	return zLogger

}

// NewGormLogger new a gorm logger.
func NewGormLogger(cfg *conf.Logger) glogger.Interface {

	zLogger := New(cfg, "gorm.log")

	level := glogger.Warn

	if cfg.Gorm.Level == "debug" {
		level = glogger.Info
	}

	logger := &zapgorm2.Logger{
		ZapLogger:                 zLogger,
		LogLevel:                  level,
		SlowThreshold:             cfg.Gorm.SlowThreshold,
		Caller:                    cfg.Gorm.Caller,
		IgnoreRecordNotFoundError: cfg.Gorm.IgnoreRecordNotFoundError,
	}

	logger.SetAsDefault() // optional: configure gorm to use this zapgorm.Logger for callbacks

	return logger
}
