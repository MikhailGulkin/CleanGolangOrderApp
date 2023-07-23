package logger

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/logger/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.SugaredLogger
}

var (
	globalLogger *Logger
	zapLogger    *zap.Logger
)

func NewLogger(loggerConfig config.LoggerConfig) Logger {
	if globalLogger == nil {
		logger := newLogger(loggerConfig)
		globalLogger = &logger
	}
	return *globalLogger
}

func newLogger(loggerConfig config.LoggerConfig) Logger {
	var zapConfig zap.Config
	if loggerConfig.Mode == "production" {
		zapConfig = zap.NewProductionConfig()
	} else if loggerConfig.Mode == "development" {
		zapConfig = zap.NewDevelopmentConfig()
		zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}
	if loggerConfig.Mode == "production" && loggerConfig.LogOutput != "" {
		zapConfig.OutputPaths = []string{loggerConfig.LogOutput}
	}

	logLevel := loggerConfig.LogLevel
	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	case "fatal":
		level = zapcore.FatalLevel
	default:
		level = zap.PanicLevel
	}
	zapConfig.Level.SetLevel(level)

	zapLogger, _ = zapConfig.Build()
	logger := newSugaredLogger(zapLogger)

	return *logger
}
func newSugaredLogger(logger *zap.Logger) *Logger {
	return &Logger{
		SugaredLogger: logger.Sugar(),
	}
}
