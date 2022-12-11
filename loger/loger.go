package loger

import (
	"Vadek/configuration"
	"context"
	"errors"
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var (
	exportUseLogger      *zap.Logger
	exportUseSugarLogger *zap.SugaredLogger
)

func NewLogger(conf *configuration.Config) *zap.Logger {

	_, err := os.Stat(conf.Vadek.LogDir)
	if err != nil {
		if os.IsNotExist(err) && !configuration.IsDev() {
			err := os.MkdirAll(conf.Vadek.LogDir, os.ModePerm)
			if err != nil {
				panic("mkdir log failed![%v]")
			}
		}
	}
	var core zapcore.Core
	if configuration.IsDev() {
		core = zapcore.NewCore(getDevEncoder(), os.Stdout, getLogLevel(conf.Log.Levels.App))
	} else {
		core = zapcore.NewCore(getProdEncoder(), getWriter(conf), zap.DebugLevel)
	}
	// 传入 zap.AddCaller() 显示打日志点的文件名和行数
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.DPanicLevel))
	exportUseLogger = logger.WithOptions(zap.AddCallerSkip(1))
	exportUseSugarLogger = exportUseLogger.Sugar()
	return logger
}

// getWriter 自定义Writer,分割日志
func getWriter(conf *configuration.Config) zapcore.WriteSyncer {
	//循环记录

	rotatingLogger := &lumberjack.Logger{
		Filename: filepath.Join(conf.Vadek.LogDir, conf.Log.FileName),
		MaxSize:  conf.Log.MaxSize,
		MaxAge:   conf.Log.MaxAge,
		Compress: conf.Log.Compress,
	}
	return zapcore.AddSync(rotatingLogger)
}

// getProdEncoder 自定义日志编码器
func getProdEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}
func getDevEncoder() zapcore.Encoder {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		panic("log level error")
	}
}

// lgoer modele
func Debugf(template string, args ...interface{}) {
	exportUseSugarLogger.Debugf(template, args...)
}

func Infof(template string, args ...interface{}) {
	exportUseSugarLogger.Infof(template, args...)
}

func Warnf(template string, args ...interface{}) {
	exportUseSugarLogger.Warnf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	exportUseSugarLogger.Errorf(template, args...)
}

func Fatalf(template string, args ...interface{}) {
	exportUseSugarLogger.Fatalf(template, args...)
}

func Debug(msg string, fields ...zap.Field) {
	exportUseLogger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	exportUseLogger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	exportUseLogger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	exportUseLogger.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	exportUseLogger.Fatal(msg, fields...)
}

func CtxDebugf(ctx context.Context, template string, args ...interface{}) {
	exportUseSugarLogger.Debugf(template, args...)
}

func CtxInfof(ctx context.Context, template string, args ...interface{}) {
	exportUseSugarLogger.Infof(template, args...)
}

func CtxWarnf(ctx context.Context, template string, args ...interface{}) {
	exportUseSugarLogger.Warnf(template, args...)
}

func CtxErrorf(ctx context.Context, template string, args ...interface{}) {
	exportUseSugarLogger.Errorf(template, args...)
}

func CtxFatalf(ctx context.Context, template string, args ...interface{}) {
	exportUseSugarLogger.Fatalf(template, args...)
}

func CtxDebug(ctx context.Context, msg string, fields ...zap.Field) {
	exportUseLogger.Debug(msg, fields...)
}

func CtxInfo(ctx context.Context, msg string, fields ...zap.Field) {
	exportUseLogger.Info(msg, fields...)
}

func CtxWarn(ctx context.Context, msg string, fields ...zap.Field) {
	exportUseLogger.Warn(msg, fields...)
}

func CtxError(ctx context.Context, msg string, fields ...zap.Field) {
	exportUseLogger.Error(msg, fields...)
}

func CtxFatal(ctx context.Context, msg string, fields ...zap.Field) {
	exportUseLogger.Fatal(msg, fields...)
}

func Sync() {
	exportUseLogger.Sync()
	exportUseSugarLogger.Sync()
}

// Gorm log
type gormLogger struct {
	logger.Config
	traceStr     string
	traceWarnStr string
	traceErrStr  string
	zapLogger    *zap.Logger
}

func NewGormLogger(conf *configuration.Config, zapLogger *zap.Logger) logger.Interface {
	logConfig := logger.Config{
		SlowThreshold:             200 * time.Millisecond,
		LogLevel:                  GetGormLogLevel(conf.Log.Levels.Gorm),
		IgnoreRecordNotFoundError: true,
		Colorful:                  configuration.IsDev(),
	}
	gl := &gormLogger{
		Config:       logConfig,
		traceStr:     "[%.3fms] [rows:%v] %s",
		traceWarnStr: "%s [%.3fms] [rows:%v] %s",
		traceErrStr:  "%s [%.3fms] [rows:%v] %s",
		zapLogger:    zapLogger,
	}
	if logConfig.Colorful {
		gl.traceStr = logger.Yellow + "[%.3fms] " + logger.BlueBold + "[rows:%v]" + logger.Reset + " %s"
		gl.traceWarnStr = "%s " + logger.Reset + logger.RedBold + "[%.3fms] " + logger.Yellow + "[rows:%v]" + logger.Magenta + " %s" + logger.Reset
		gl.traceErrStr = logger.MagentaBold + "%s " + logger.Reset + logger.Yellow + "[%.3fms] " + logger.BlueBold + "[rows:%v]" + logger.Reset + " %s"
	}
	return gl
}

func (l *gormLogger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *l
	newLogger.LogLevel = level
	return &newLogger
}

const level = 2

func (l *gormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	l.zapLogger.WithOptions(zap.AddCallerSkip(getCallerSkip()-level)).Sugar().Infof(msg, data...)
}

func (l *gormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	l.zapLogger.WithOptions(zap.AddCallerSkip(getCallerSkip()-level)).Sugar().Warnf(msg, data...)
}

func (l *gormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	l.zapLogger.WithOptions(zap.AddCallerSkip(getCallerSkip()-level)).Sugar().Errorf(msg, data...)
}

func (l *gormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if l.LogLevel <= logger.Silent {
		return
	}

	elapsed := time.Since(begin)
	switch {
	case err != nil && l.LogLevel >= logger.Error && (!errors.Is(err, gorm.ErrRecordNotFound) || !l.IgnoreRecordNotFoundError):
		sql, rows := fc()
		if rows == -1 {
			l.zapLogger.WithOptions(zap.AddCallerSkip(getCallerSkip()-level)).Sugar().Errorf(l.traceErrStr, err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.zapLogger.WithOptions(zap.AddCallerSkip(getCallerSkip()-level)).Sugar().Errorf(l.traceErrStr, err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= logger.Warn:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
		if rows == -1 {
			l.zapLogger.WithOptions(zap.AddCallerSkip(getCallerSkip()-level)).Sugar().Warnf(l.traceWarnStr, slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.zapLogger.WithOptions(zap.AddCallerSkip(getCallerSkip()-level)).Sugar().Warnf(l.traceWarnStr, slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case l.LogLevel == logger.Info:
		sql, rows := fc()
		if rows == -1 {
			l.zapLogger.WithOptions(zap.AddCallerSkip(getCallerSkip()-level)).Sugar().Infof(l.traceStr, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.zapLogger.WithOptions(zap.AddCallerSkip(getCallerSkip()-level)).Sugar().Infof(l.traceStr, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	}
}

func getCallerSkip() int {
	for i := 3; i < 15; i++ {
		pc := make([]uintptr, 1)
		numFrames := runtime.Callers(i, pc)
		if numFrames < 1 {
			return i
		}
		frame, _ := runtime.CallersFrames(pc).Next()
		if !strings.Contains(frame.Function, "gorm.io") && !strings.Contains(frame.Function, "github.com/go-sonic/sonic/dal") {
			return i
		}
	}
	return 0
}

func GetGormLogLevel(level string) logger.LogLevel {
	switch level {
	case "info":
		return logger.Info
	case "warn":
		return logger.Warn
	case "error":
		return logger.Error
	case "silent":
		return logger.Silent
	default:
		panic("log level error")
	}
}
