package log

import (
	"errors"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	ErrInvalidLogLevel     = "invalid logging level"
	ErrCannotCreateLogFile = "cannot create log file"
)

var stringLevelToZapcore map[string]zapcore.Level = map[string]zapcore.Level{
	ErrorLevel: zapcore.ErrorLevel,
	DebugLevel: zapcore.DebugLevel,
	InfoLevel:  zapcore.InfoLevel,
	WarnLevel:  zapcore.WarnLevel,
}

func convertLevelToZapcore(level string) (zapcore.Level, error) {
	v, ok := stringLevelToZapcore[level]
	if !ok {
		return zapcore.Level(1), errors.New(ErrInvalidLogLevel)
	}

	return v, nil
}

func getStaticEnabler(enabled bool) zapcore.LevelEnabler {
	return zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return enabled
	})
}

func InitZapLogger(config *Config) (*zap.Logger, error) {
	logLevel, err := convertLevelToZapcore(config.Level)
	if err != nil {
		return nil, err
	}

	consoleEnabler := getStaticEnabler(config.Console)
	consoleSyncer := zapcore.Lock(os.Stdout)
	consoleEncoder := zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig())

	cfg := zap.NewProductionConfig()
	cfg.Encoding = "json"
	cfg.OutputPaths = []string{config.File}
	cfg.Level = zap.NewAtomicLevelAt(logLevel)

	logger, err := cfg.Build(zap.WrapCore(func(c zapcore.Core) zapcore.Core {
		return zapcore.NewTee(
			c,
			zapcore.NewCore(consoleEncoder, consoleSyncer, consoleEnabler),
		)
	}))
	if err != nil {
		return nil, err
	}

	zap.ReplaceGlobals(logger)

	return logger, nil
}
