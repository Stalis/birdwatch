package log

import (
	"os"

	"github.com/pkg/errors"
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

func getConditionalEnabler(enabled bool, level zapcore.Level) zapcore.LevelEnabler {
	return zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return enabled && l >= level
	})
}

func InitZapLogger(config *Config) (*zap.Logger, error) {
	logLevel, err := convertLevelToZapcore(config.Level)
	if err != nil {
		return nil, err
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	file, err := os.OpenFile(config.File, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return nil, errors.Wrap(err, ErrCannotCreateLogFile)
	}

	fileEnabler := getConditionalEnabler(true, logLevel)
	fileSyncer := zapcore.AddSync(file)
	fileEncoder := zapcore.NewJSONEncoder(encoderConfig)

	consoleEncoderConfig := encoderConfig
	consoleEncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006 Jan 1 15:04:05.000")
	consoleEnabler := getConditionalEnabler(config.Console, logLevel)
	consoleSyncer := zapcore.Lock(os.Stdout)
	consoleEncoder := zapcore.NewConsoleEncoder(consoleEncoderConfig)

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, fileSyncer, fileEnabler),
		zapcore.NewCore(consoleEncoder, consoleSyncer, consoleEnabler),
	)

	logger := zap.New(core)

	zap.ReplaceGlobals(logger)

	return logger, nil
}
