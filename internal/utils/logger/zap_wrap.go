package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(name string) *zap.SugaredLogger {

	var err error

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	logLevelEnv := os.Getenv("NOTELS_LOG_LEVEL")
	err = config.Level.UnmarshalText([]byte(logLevelEnv))
	if err != nil {
		panic(err)
	}

	var logger *zap.Logger

	logger, err = config.Build()
	if err != nil {
		panic(err)
	}

	return logger.Named(name).Sugar()
}
