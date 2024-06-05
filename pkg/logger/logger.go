package logger

import (
	"io"
	"os"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Field struct {
	Name string
	Data interface{}
}

var logger *zap.SugaredLogger

func Init(mode string) {
	var level zap.LevelEnablerFunc
	encoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		TimeKey:       "created_at",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "from",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.CapitalLevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})

	switch mode {
	case "debug":
		level = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zapcore.DebugLevel
		})
	case "info":
		level = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zapcore.InfoLevel
		})
	case "warn":
		level = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zapcore.WarnLevel
		})
	case "error":
		level = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zapcore.ErrorLevel
		})
	}

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), level),
	)
	log := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	logger = log.Sugar()
}
func getWriter(filename string) io.Writer {
	hook, err := rotatelogs.New(
		strings.Replace(filename, ".log", "", -1)+"-%Y%m%d%H.log",
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationTime(time.Hour),
	)
	if err != nil {
		panic(err)
	}
	return hook
}
func Debug(args ...interface{}) {
	logger.Debug(args...)
}
func Debugf(template string, args ...interface{}) {
	logger.Debugf(template, args...)
}
func Info(args ...interface{}) {
	logger.Info(args...)
}
func Infof(template string, args ...interface{}) {
	logger.Infof(template, args...)
}
func Warn(args ...interface{}) {
	logger.Warn(args...)
}
func Warnf(template string, args ...interface{}) {
	logger.Warnf(template, args...)
}
func Error(args ...interface{}) {
	logger.Error(args...)
}
func Errorf(template string, args ...interface{}) {
	logger.Errorf(template, args...)
}
func DPanic(args ...interface{}) {
	logger.DPanic(args...)
}
func DPanicf(template string, args ...interface{}) {
	logger.DPanicf(template, args...)
}
func Panic(args ...interface{}) {
	logger.Panic(args...)
}
func Panicf(template string, args ...interface{}) {
	logger.Panicf(template, args...)
}
func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}
func Fatalf(template string, args ...interface{}) {
	logger.Fatalf(template, args...)
}
