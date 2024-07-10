package log

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type contextKey string

const logKey contextKey = "log"

func Init() (*zap.Logger, *zap.SugaredLogger, error) {
	logConfig := zap.NewDevelopmentConfig()
	logConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logConfig.DisableCaller = true
	logConfig.Level = zap.NewAtomicLevelAt(zap.InfoLevel)

	logger, err := logConfig.Build()
	if err != nil {
		return nil, nil, fmt.Errorf("error building logger: %w", err)
	}

	sugar := logger.Sugar()

	return logger, sugar, nil
}

func InjectIntoContext(ctx context.Context, log *zap.SugaredLogger) context.Context {
	return context.WithValue(ctx, logKey, log)
}

func FromContext(ctx context.Context) *zap.SugaredLogger {
	logValue := ctx.Value(logKey)
	if logValue == nil {
		return nil
	}

	log, ok := logValue.(*zap.SugaredLogger)
	if !ok {
		return nil
	}

	return log
}
