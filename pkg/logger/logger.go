package logger

import (
	"go.uber.org/zap"
)

// InitLogger initializes Zap logger
func InitLogger() *zap.Logger {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	return logger
}
