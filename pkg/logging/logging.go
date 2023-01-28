package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var ee *zap.Logger

type Logger struct {
	*zap.Logger
}

var logLevel = map[string]zapcore.Level{
	"DEBUG": zapcore.DebugLevel,
	"INFO":  zapcore.InfoLevel,
	"ERROR": zapcore.ErrorLevel,
	"WARN":  zapcore.WarnLevel,
}

func GetLogger() Logger {
	return Logger{ee}
}

func init() {
	config := zap.NewProductionConfig()
	level, ok := logLevel[os.Getenv("VMWARE_EXPORTER_LOG_LEVEL")]
	if ok {
		config.Level.SetLevel(level)
	}
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.MessageKey = "message"
	ee, _ = config.Build()
}

func (l *Logger) String(key, val string) zap.Field {
	return zap.String(key, val)
}

func (l *Logger) Int(key string, val int) zap.Field {
	return zap.Int(key, val)
}

func (l *Logger) Duration(key string, time time.Duration) zap.Field {
	return zap.Duration(key, time)
}

func (l *Logger) Error(msg string) {
	l.Logger.Sugar().Error(msg)
}

func (l *Logger) Infof(msg string, args ...interface{}) {
	l.Logger.Sugar().Infof(msg, args...)
}

func (l *Logger) Debugf(msg string, args ...interface{}) {
	l.Logger.Sugar().Debugf(msg, args...)
}

func (l *Logger) GetLevel() string {
	return os.Getenv("APP_LOG_LEVEL")
}
