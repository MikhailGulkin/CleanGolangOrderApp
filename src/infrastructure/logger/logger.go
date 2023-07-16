package logger

import (
	"fmt"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/logger/config"
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

func GetLogger(loggerConfig config.LoggerConfig) Logger {
	if globalLogger == nil {
		logger := newLogger(loggerConfig)
		globalLogger = &logger
	}
	return *globalLogger
}
func newLogger(loggerConfig config.LoggerConfig) Logger {
	zapConfig := zap.NewDevelopmentConfig()
	logOutput := loggerConfig.LogOutput

	if loggerConfig.Mode == "development" {
		fmt.Println("encode level")
		zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	if loggerConfig.Mode == "production" && logOutput != "" {
		zapConfig.OutputPaths = []string{logOutput}
	}

	logLevel := loggerConfig.LogLevel
	level := zap.PanicLevel
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
