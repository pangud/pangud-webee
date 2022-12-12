package zapgorm2

import (
	"context"
	"errors"
	"fmt"
	"path/filepath"
	"runtime"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

const timeFormatLayout = "2006-01-02 15:04:05.000"

type Logger struct {
	ZapLogger                 *zap.Logger
	LogLevel                  gormlogger.LogLevel
	SlowThreshold             time.Duration
	Caller                    bool
	IgnoreRecordNotFoundError bool
}

func New(zapLogger *zap.Logger) Logger {
	return Logger{
		ZapLogger:                 zapLogger,
		LogLevel:                  gormlogger.Warn,
		SlowThreshold:             100 * time.Millisecond,
		Caller:                    false,
		IgnoreRecordNotFoundError: false,
	}
}

func (l Logger) SetAsDefault() {
	gormlogger.Default = l
}

func (l Logger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	return Logger{
		ZapLogger:                 l.ZapLogger,
		SlowThreshold:             l.SlowThreshold,
		LogLevel:                  level,
		Caller:                    l.Caller,
		IgnoreRecordNotFoundError: l.IgnoreRecordNotFoundError,
	}
}

func (l Logger) Info(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel < gormlogger.Info {
		return
	}
	l.logger().Sugar().Debugf(str, args...)
}

func (l Logger) Warn(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel < gormlogger.Warn {
		return
	}
	l.logger().Sugar().Warnf(str, args...)
}

func (l Logger) Error(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel < gormlogger.Error {
		return
	}
	l.logger().Sugar().Errorf(str, args...)
}

func (l Logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= 0 {
		return
	}
	elapsed := time.Since(begin)
	switch {
	case err != nil && l.LogLevel >= gormlogger.Error && (!l.IgnoreRecordNotFoundError || !errors.Is(err, gorm.ErrRecordNotFound)):
		sql, rows := fc()

		var fields = make([]zap.Field, 0)
		fields = append(fields, zap.String("ts", time.Now().Format(timeFormatLayout)))
		if l.Caller {
			fields = append(fields, zap.String("caller", printCaller()))
		}
		fields = append(fields, zap.Error(err))
		fields = append(fields, zap.Duration("elapsed", elapsed))
		fields = append(fields, zap.Int64("rows", rows))
		fields = append(fields, zap.String("sql", sql))
		l.logger().Error("trace", fields...)

	case l.SlowThreshold != 0 && elapsed > l.SlowThreshold && l.LogLevel >= gormlogger.Warn:
		sql, rows := fc()

		var fields = make([]zap.Field, 0)
		fields = append(fields, zap.String("ts", time.Now().Format(timeFormatLayout)))
		if l.Caller {
			fields = append(fields, zap.String("caller", printCaller()))
		}
		fields = append(fields, zap.Duration("elapsed", elapsed))
		fields = append(fields, zap.Int64("rows", rows))
		fields = append(fields, zap.String("sql", sql))
		l.logger().Warn("trace", fields...)
	case l.LogLevel >= gormlogger.Info:
		sql, rows := fc()

		var fields = make([]zap.Field, 0)
		fields = append(fields, zap.String("ts", time.Now().Format(timeFormatLayout)))
		if l.Caller {
			fields = append(fields, zap.String("caller", printCaller()))
		}
		fields = append(fields, zap.Duration("elapsed", elapsed))
		fields = append(fields, zap.Int64("rows", rows))
		fields = append(fields, zap.String("sql", sql))
		l.logger().Debug("trace", fields...)
	}
}
func printCaller() string {
	//add caller
	_, filePath, lineNum, ok := runtime.Caller(4)
	if !ok {
		_, filePath, lineNum, ok = runtime.Caller(3)
	}
	path := filepath.Base(filepath.Join(filePath, filepath.ToSlash("../")))
	filePath = filepath.Join(path, filepath.Base(filePath))
	return filePath + fmt.Sprintf(":%v ", lineNum)
}

//var (
//	gormPackage    = filepath.Join("gorm.io", "gorm")
//	zapgormPackage = filepath.Join("moul.io", "zapgorm2")
//)

func (l Logger) logger() *zap.Logger {
	//for i := 2; i < 15; i++ {
	//	_, file, _, ok := runtime.Caller(i)
	//	switch {
	//	case !ok:
	//	case strings.HasSuffix(file, "_test.go"):
	//	case strings.Contains(file, gormPackage):
	//	case strings.Contains(file, zapgormPackage):
	//	default:
	//		return l.ZapLogger.WithOptions(zap.AddCallerSkip(i))
	//	}
	//}
	return l.ZapLogger
}
