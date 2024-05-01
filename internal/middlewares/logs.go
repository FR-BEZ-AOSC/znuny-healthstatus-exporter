package middlewares

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(path string) *zap.Logger {

	localLogger, _ := zap.NewProduction()
	defer localLogger.Sync()
	
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		TemporaryLogger("Failed to initialize the logger", err)
	}

	cfg := zap.NewProductionConfig()
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg.EncoderConfig),
		zapcore.NewMultiWriteSyncer(os.Stdout, file),
		zapcore.InfoLevel,
	)

	logger := zap.New(core)
	logger.Sync()

	return logger
}

func TemporaryLogger(msg string, err error)  {
	localLogger, _ := zap.NewProduction()
	defer localLogger.Sync()

	localLogger.Panic(
		msg,
		zap.Error(err),
	)	
}